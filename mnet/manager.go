package mnet

import (
	"code.google.com/p/goprotobuf/proto"
	"log"
	"net"
	"runtime"
	"strings"
)

type ConnectionHandler interface {
	OnValidateToken(token string) int64 // token to userId
	OnConnected(userId int64)
	OnDisconnected(userId int64)
	OnRecievePayload(userId int64, payload *Payload)
}

func NewManager() *Manager {
	return &Manager{server: NewServer()}
}

// implements Pushing inteface
type Manager struct {
	Handler      ConnectionHandler
	server       *Server
	nextClientId int64
}

func (p *Manager) NewChannel(subsId []int64) (channelId int64, err error) {
	channel := NewChannel(subsId)
	return channel.Id, nil
}

func (p *Manager) PushToUser(userId int64, message *Message) (err error) {
	// packet := &Packet{Body: message}
	var body []byte
	body, err = proto.Marshal(message.MSG)
	if err != nil {
		return
	}
	payload := &Payload{Code: uint16(message.Code), Body: body}
	client := p.server.clients[userId]
	client.Write(payload)

	err = nil
	return
}

func (p *Manager) PushToChannel(chanelId int64, message *Message) (err error) {
	channel := GetChannel(chanelId)
	if channel != nil {
		for _, v := range channel.Subs {
			p.PushToUser(v, message)
		}
	}

	err = nil
	return
}

func (p *Manager) HandleConnection(handler ConnectionHandler) {
	p.Handler = handler
}

func (p *Manager) Serve(listener net.Listener) {
	log.Printf("TCP: listening on %s", listener.Addr().String())

	for {
		conn, err := listener.Accept()
		if err != nil {
			if nerr, ok := err.(net.Error); ok && nerr.Temporary() {
				log.Printf("NOTICE: temporary Accept() failure - %s", err.Error())
				runtime.Gosched()
				continue
			}

			// theres no direct way to detect this error because it is not exposed
			if !strings.Contains(err.Error(), "use of closed network connection") {
				log.Printf("ERROR: listener.Accept() - %s", err.Error())
			}
			break
		}

		go p.handleTcpClient(conn)
	}

	log.Printf("TCP: closing %s", listener.Addr().String())
}

func (p *Manager) handleTcpClient(conn net.Conn) {
	log.Printf("New TCP Client %s", "...")
	defer func() {
		err := conn.Close()
		if err != nil {
			// s.errCh <- err
			log.Printf("Client close error: %s", err.Error())
		}
		log.Print("Client disconnected")
		if p.Handler != nil {
			//p.Handler.OnDisconnected(userId)
		}
	}()

	client := NewClient(conn, p.server, p)
	client.id = p.nextClientId
	p.nextClientId++
	p.server.Add(client)
	if p.Handler != nil {
		// p.Handler.OnConnected(userId)
	}

	// payload := Payload{Code: 1, Body: []byte("hello")}

	// client.Write(&payload)
	client.Listen()
}
