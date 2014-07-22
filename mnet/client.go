package mnet

import (
	"bytes"
	// "encoding/json"
	"fmt"
	"io"
	"log"
	"net"
)

const channelBufSize = 100

type Client struct {
	id      int64
	conn    net.Conn
	server  *Server
	manager *Manager
	ch      chan *Payload
	doneCh  chan bool
	buf     bytes.Buffer
}

// Create new chat client.
func NewClient(conn net.Conn, server *Server, manager *Manager) *Client {

	if conn == nil {
		panic("conn cannot be nil")
	}

	if server == nil {
		panic("server cannot be nil")
	}

	ch := make(chan *Payload, channelBufSize)
	doneCh := make(chan bool)

	return &Client{conn: conn, server: server, ch: ch, doneCh: doneCh, manager: manager}
}

func (c *Client) Conn() net.Conn {
	return c.conn
}

func (c *Client) Write(payload *Payload) {
	select {
	case c.ch <- payload:
	default:
		c.server.Del(c)
		err := fmt.Errorf("Client(%d) is disconnected.", c.id)
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
		case payload := <-c.ch:
			log.Printf("Client(%d) send payload: %#v", c.id, payload)
			b, _ := payload.Encode()
			c.conn.Write(b)

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
			b := make([]byte, 10)
			length, err := c.conn.Read(b)
			if err == io.EOF {
				c.doneCh <- true
			} else if err != nil {
				c.server.Err(err)
			} else {
				log.Printf("Client(%d) recv raw: % x\n", c.id, b)
				c.buf.Write(b[0:length])

				payload := &Payload{}
				err, more, left := payload.Decode(c.buf.Bytes())
				if err != nil {
					log.Printf("Client(%d) payload.Decode err: %s", c.id, err.Error())
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

					c.manager.Handler.OnRecievePayload(c.id, payload)
				}
			}
		}
	}
}
