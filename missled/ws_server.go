package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"github.com/miraclew/mrs/push"
	"log"
	"net"
	"net/http"
	"strings"
)

func init() {
}

func wsServe(listener net.Listener) {
	log.Printf("WS: listening on %s", listener.Addr().String())

	var maxId int64 = 0
	s := push.NewServer()
	onConnected := func(ws *websocket.Conn) {
		defer func() {
			err := ws.Close()
			if err != nil {
				// s.errCh <- err
				fmt.Println(err)
			}
		}()

		maxId++
		client := push.NewClient(maxId, ws, s)
		s.Add(client)
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
