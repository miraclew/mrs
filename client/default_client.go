package main

import (
	"bytes"
	"log"
	"net"
	// "fmt"
	"code.google.com/p/goprotobuf/proto"
	"github.com/miraclew/mrs/mnet"
	"github.com/miraclew/mrs/pb"
	"io"
)

const (
	STATE_DISCONNECTED = 1
	STATE_CONNECTED    = 2
	STATE_READY        = 3
	STATE_GAMING       = 4
)

type DefaultClient struct {
	user  *User
	conn  *net.TCPConn
	state int
	buf   bytes.Buffer
}

func (d *DefaultClient) runAs(user *User) {
	d.user = user
	d.state = STATE_DISCONNECTED
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8081")
	d.handleErr(err)

	// connect
	conn, err := net.DialTCP("tcp", nil, addr)
	d.handleErr(err)
	d.conn = conn
	d.state = STATE_CONNECTED

	d.auth()

	for {
		b := make([]byte, 10)
		length, err := d.conn.Read(b)
		if err == io.EOF {
			d.handleErr(err)
		} else if err != nil {
			d.handleErr(err)
		} else {
			d.buf.Write(b[0:length])

			payload := &mnet.Payload{}
			err, more, left := payload.Decode(d.buf.Bytes())
			if err != nil {
				log.Printf("User(%d) payload.Decode err: %s", d.user.Id, err.Error())
				return
			}

			if more {
				continue
			} else {
				d.buf.Reset()
				if len(left) > 0 {
					d.buf.Write(left)
				}

				d.OnRecievePayload(payload)
				log.Printf("User(%d) receive payload: %#v\n", d.user.Id, payload)
			}
		}
	}
}

func (d *DefaultClient) auth() {
	auth := &pb.CAuth{}
	auth.UserName = &d.user.UserName
	auth.Password = &d.user.Password
	d.send(pb.Code_C_AUTH, auth)
}

func (d *DefaultClient) send(code pb.Code, msg proto.Message) (err error) {
	var body []byte
	body, err = proto.Marshal(msg)
	if err != nil {
		return
	}
	payload := &mnet.Payload{Code: uint16(code), Body: body}

	b, _ := payload.Encode()
	d.conn.Write(b)
	return
}

func (d *DefaultClient) OnRecievePayload(payload *mnet.Payload) (err error) {
	switch pb.Code(payload.Code) {
	case pb.Code_E_AUTH:
		eauth := &pb.EAuth{}
		err = proto.Unmarshal(payload.Body, eauth)
		if err == nil {
			log.Printf("EAuth: %s", eauth.String())
		}
	case pb.Code_E_MATCH_INIT:
	case pb.Code_E_MATCH_TURN:
	case pb.Code_E_MATCH_END:
	case pb.Code_E_PLAYER_MOVE:
	case pb.Code_E_PLAYER_FIRE:
	case pb.Code_E_PLAYER_HIT:

	}

	return
}

func (d *DefaultClient) handleErr(err error) {
	log.Fatalf("User(%d) error: %s\n", d.user.Id, err.Error())
}
