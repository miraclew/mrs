package mnet

import (
	"code.google.com/p/goprotobuf/proto"
	"log"
	"net"
	"runtime"
	"strings"
)

type ConnectionHandler interface {
	OnConnected(clientId int64)
	OnDisconnected(clientId int64)
	OnRecievePayload(clientId int64, payload *Payload)
}

func NewManager(server *Server) *Manager {
	return &Manager{server: server, nextClientId: 1}
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

func (p *Manager) PushToClient(clientId int64, message *Message) (err error) {
	log.Printf("PushToClient(%d): %s", clientId, message.String())
	// packet := &Packet{Body: message}
	var body []byte
	body, err = proto.Marshal(message.MSG)
	if err != nil {
		return
	}
	payload := &Payload{Code: uint16(message.Code), Body: body}
	client := p.server.clients[clientId]
	client.Write(payload)

	err = nil
	return
}

func (p *Manager) PushToChannel(chanelId int64, message *Message) (err error) {
	channel := GetChannel(chanelId)
	if channel != nil {
		for _, v := range channel.Subs {
			p.PushToClient(v, message)
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
	defer func() {
		err := conn.Close()
		if err != nil {
			// s.errCh <- err
			log.Printf("Client close error: %s", err.Error())
		}
		log.Print("Client disconnected")
		if p.Handler != nil {
			//p.Handler.OnDisconnected(clientId)
		}
	}()

	client := NewClient(conn, p.server, p)
	client.id = p.nextClientId
	p.nextClientId++
	p.server.Add(client)
	if p.Handler != nil {
		// p.Handler.OnConnected(clientId)
	}

	log.Printf("New Client: %d", client.id)
	client.Listen()
}
