package main

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

func init() {
}

func wsServe(listener net.Listener) {
	log.Printf("WS: listening on %s", listener.Addr().String())

	onConnected := func(ws *websocket.Conn) {
		defer func() {
			err := ws.Close()
			if err != nil {
				// s.errCh <- err
			}
		}()

		// client := NewClient(ws, s)
		// s.Add(client)
		// client.Listen()
		fmt.Println("new client")
	}

	handler := websocket.Handler(onConnected)
	server := &http.Server{
		Handler: handler,
	}

	err := server.Serve(listener)
	// theres no direct way to detect this error because it is not exposed
	if err != nil && !strings.Contains(err.Error(), "use of closed network connection") {
		log.Printf("ERROR: ws.Serve() - %s", err.Error())
	}

	log.Printf("HTTP: closing %s", listener.Addr().String())
}
