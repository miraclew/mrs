package push

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

type ConnectionHandler interface {
	OnValidateToken(token string) int64 // token to userId
	OnConnected(userId int64)
	OnDisconnected(userId int64)
}

// implements Pushing inteface
type Pusher struct {
	handler ConnectionHandler
	server  Server
}

func (p *Pusher) NewChannel(subsId []int64) (channelId int64, err error) {
	channel := NewChannel(subsId)
	return channel.Id, nil
}

func (*Pusher) PushToUser(userId int64, message interface{}) (err error) {
	packet := &Packet{Body: message}
	client := server.clients[userId]
	client.Write(packet)

	err = nil
	return
}

func (p *Pusher) PushToChannel(chanelId int64, message interface{}) (err error) {
	channel := GetChannel(chanelId)
	if channel != nil {
		for _, v := range channel.Subs {
			p.PushToUser(v, message)
		}
	}

	err = nil
	return
}

func (p *Pusher) Serve(listener net.Listener) {
	log.Printf("WS: listening on %s", listener.Addr().String())

	s := NewServer()
	onConnected := func(ws *websocket.Conn) {
		var userId int64 = 0
		defer func() {
			err := ws.Close()
			if err != nil {
				// s.errCh <- err
				fmt.Println(err)
			}
			if p.handler != nil && userId != 0 {
				p.handler.OnDisconnected(userId)
			}
		}()

		token := ws.Request().URL.Query().Get("token")
		userId = p.handler.OnValidateToken(token)
		log.Printf("New connection, token->userId %s -> %d \n", token, userId)
		if userId == 0 {
			return
		}
		client := NewClient(userId, ws, s)
		s.Add(client)
		if p.handler != nil {
			p.handler.OnConnected(userId)
		}
		client.Listen()
	}

	go s.Listen()

	wsHandler := &websocket.Server{Handler: websocket.Handler(onConnected)}

	httpServer := &http.Server{
		Handler: wsHandler,
	}

	err := httpServer.Serve(listener)
	// theres no direct way to detect this error because it is not exposed
	if err != nil && !strings.Contains(err.Error(), "use of closed network connection") {
		log.Printf("ERROR: ws.Serve() - %s", err.Error())
	}

	log.Printf("HTTP: closing %s", listener.Addr().String())
}

func (p *Pusher) HandleConnection(handler ConnectionHandler) {
	p.handler = handler
}
