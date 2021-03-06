package push

import (
	"log"
)

type Server struct {
	// messages  []*Message
	clients map[int64]*Client
	addCh   chan *Client
	delCh   chan *Client
	// sendAllCh chan *Message
	doneCh chan bool
	errCh  chan error
}

var server *Server

func NewServer() *Server {
	if server == nil {
		// messages := []*Message{}
		clients := make(map[int64]*Client)
		addCh := make(chan *Client)
		delCh := make(chan *Client)
		// sendAllCh := make(chan *Message)
		doneCh := make(chan bool)
		errCh := make(chan error)

		server = &Server{
			// messages,
			clients,
			addCh,
			delCh,
			// sendAllCh,
			doneCh,
			errCh,
		}
	}

	return server
}

func (s *Server) Add(c *Client) {
	s.addCh <- c
}

func (s *Server) Del(c *Client) {
	s.delCh <- c
}

// func (s *Server) SendAll(msg *Message) {
// 	s.sendAllCh <- msg
// }

func (s *Server) Done() {
	s.doneCh <- true
}

func (s *Server) Err(err error) {
	s.errCh <- err
}

// func (s *Server) sendPastMessages(c *Client) {
// 	for _, msg := range s.messages {
// 		c.Write(msg)
// 	}
// }

// func (s *Server) sendAll(msg *Message) {
// 	for _, c := range s.clients {
// 		c.Write(msg)
// 	}
// }

// Listen and serve.
// It serves client connection and broadcast request.
func (s *Server) Listen() {

	log.Println("Listening server...")

	for {
		select {

		// Add new a client
		case c := <-s.addCh:
			log.Println("Added new client")
			s.clients[c.id] = c
			log.Println("Now", len(s.clients), "clients connected.")
			// s.sendPastMessages(c)

		// del a client
		case c := <-s.delCh:
			log.Println("Delete client")
			delete(s.clients, c.id)

		// broadcast message for all clients
		// case msg := <-s.sendAllCh:
		// 	log.Println("Send all:", msg)
		// 	s.messages = append(s.messages, msg)
		// 	s.sendAll(msg)

		case err := <-s.errCh:
			log.Println("Error:", err.Error())

		case <-s.doneCh:
			return
		}
	}
}
