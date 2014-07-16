package mnet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
)

const channelBufSize = 100

type Client struct {
	id     int64
	conn   net.Conn
	server *Server
	ch     chan *Packet
	doneCh chan bool
	buf    bytes.Buffer
}

// Create new chat client.
func NewClient(conn net.Conn, server *Server) *Client {

	if conn == nil {
		panic("conn cannot be nil")
	}

	if server == nil {
		panic("server cannot be nil")
	}

	ch := make(chan *Packet, channelBufSize)
	doneCh := make(chan bool)

	return &Client{conn: conn, server: server, ch: ch, doneCh: doneCh}
}

func (c *Client) Conn() net.Conn {
	return c.conn
}

func (c *Client) Write(msg *Packet) {
	select {
	case c.ch <- msg:
	default:
		c.server.Del(c)
		err := fmt.Errorf("client %d is disconnected.", c.id)
		c.server.Err(err)
	}
}

func (c *Client) Done() {
	c.doneCh <- true
}

// Listen Write and Read request via chanel
func (c *Client) Listen() {
	go c.listenWrite()
	c.listenRead()
}

// Listen write request via chanel
func (c *Client) listenWrite() {
	log.Println("Listening write to client")
	for {
		select {

		// send message to the client
		case msg := <-c.ch:
			log.Println("Send:", msg)
			bytes, _ := json.Marshal(msg)
			c.conn.Write(bytes)

		// receive done request
		case <-c.doneCh:
			c.server.Del(c)
			c.doneCh <- true // for listenRead method
			return
		}
	}
}

// Listen read request via chanel
func (c *Client) listenRead() {
	log.Println("Listening read from client")
	for {
		select {

		// receive done request
		case <-c.doneCh:
			c.server.Del(c)
			c.doneCh <- true // for listenWrite method
			return

		// read data from net connection
		default:
			// var msg Packet
			b := make([]byte, 10)
			length, err := c.conn.Read(b)
			if err == io.EOF {
				c.doneCh <- true
			} else if err != nil {
				c.server.Err(err)
			} else {
				c.buf.Write(b[0:length])

				payload := Payload{}
				err, more, left := payload.Decode(c.buf.Bytes())
				if err != nil {
					fmt.Printf("payload.Decode err: %s", err.Error())
					c.doneCh <- true
					return
				}

				if more {
					continue
				} else {
					c.buf.Reset()
					if len(left) > 0 {
						c.buf.Write(left)
					}

					fmt.Printf("Receive msg: %#v\n", payload)
				}
			}
		}
	}
}
