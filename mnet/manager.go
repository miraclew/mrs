package mnet

import (
	"fmt"
	"log"
	"net"
	"runtime"
	"strings"
)

type ConnectionHandler interface {
	OnValidateToken(token string) int64 // token to userId
	OnConnected(userId int64)
	OnDisconnected(userId int64)
}

func NewManager(server *Server, handler ConnectionHandler) *Manager {
	return &Manager{handler, server}
}

// implements Pushing inteface
type Manager struct {
	handler ConnectionHandler
	server  *Server
}

func (p *Manager) NewChannel(subsId []int64) (channelId int64, err error) {
	channel := NewChannel(subsId)
	return channel.Id, nil
}

func (p *Manager) PushToUser(userId int64, message interface{}) (err error) {
	packet := &Packet{Body: message}
	client := p.server.clients[userId]
	client.Write(packet)

	err = nil
	return
}

func (p *Manager) PushToChannel(chanelId int64, message interface{}) (err error) {
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
	p.handler = handler
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
			fmt.Println(err)
		}
		if p.handler != nil {
			//p.handler.OnDisconnected(userId)
		}
	}()

	client := NewClient(conn, p.server)
	if p.handler != nil {
		// p.handler.OnConnected(userId)
	}
	client.Listen()
}
