package push

import (
	"code.google.com/p/go.net/websocket"
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func TestA(t *testing.T) {
	origin := "http://localhost/"
	url := "ws://localhost:8081/ws?token=aaaaabbb"
	ws, err := websocket.Dial(url, "", origin)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("connected")
	msg := "hello from client"

	if err := websocket.JSON.Send(ws, msg); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Send: %#v \n", msg)

	for {
		if err = websocket.JSON.Receive(ws, &msg); err != nil {
			log.Fatal(err)
		}

		bytes, _ := json.Marshal(msg)
		fmt.Printf("Received: %s \n", string(bytes))
	}
}
