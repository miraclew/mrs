package main

import (
	"bytes"
	"log"
	"net"
	// "fmt"
	"code.google.com/p/goprotobuf/proto"
	"github.com/miraclew/mrs/missle/model"
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
	user  model.User
	conn  *net.TCPConn
	state int
	buf   bytes.Buffer
	match *pb.EMatcInit
}

func (d *DefaultClient) run() {
	log.Printf("%s born...\n", d.user.NickName)
	d.state = STATE_DISCONNECTED
	addr, err := net.ResolveTCPAddr("tcp", "127.0.0.1:8081")
	d.handleErr(err)

	// connect
	log.Printf("%s connecting...\n", d.user.NickName)
	conn, err := net.DialTCP("tcp", nil, addr)
	d.handleErr(err)
	d.conn = conn
	d.state = STATE_CONNECTED

	log.Printf("%s auth...\n", d.user.NickName)
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
	// log.Printf("%s receive payload: %#v", d.user.NickName, payload)
	switch pb.Code(payload.Code) {
	case pb.Code_E_AUTH:
		eauth := &pb.EAuth{}
		err = proto.Unmarshal(payload.Body, eauth)
		if err == nil {
			if *eauth.Code == 0 {
				log.Println("Auth OK, enter game")
				d.send(pb.Code_C_MATCH_ENTER, &pb.CMatchEnter{})
			} else {
				log.Println("Auth Fail")
			}
		}
	case pb.Code_E_MATCH_INIT:
		mi := &pb.EMatcInit{}
		err = proto.Unmarshal(payload.Body, mi)
		if err == nil {
			d.match = mi
			log.Printf("%s EMatchInit: %#v\n", d.user.NickName, mi.String())
		}

	case pb.Code_E_MATCH_TURN:
		mt := &pb.EMatchTurn{}
		err = proto.Unmarshal(payload.Body, mt)
		if err == nil {
			log.Printf("%s EMatchTurn: %#v\n", d.user.NickName, mt.String())
			// fire
			matchId := d.match.GetMatchId()
			fire := &pb.CPlayerFire{}
			fire.MatchId = &matchId
			fire.PlayerId = &d.user.Id
			var x float32 = 0.1
			var y float32 = 0.2
			fire.Velocity = &pb.Point{X: &x, Y: &y}
			d.send(pb.Code_C_PLAYER_FIRE, fire)
		}
	case pb.Code_E_MATCH_END:
		me := &pb.EMatchEnd{}
		err = proto.Unmarshal(payload.Body, me)
		if err == nil {
			log.Printf("%s Game Over: %s", d.user.NickName, me.String())
			d.match = nil
		}
	case pb.Code_E_PLAYER_MOVE:
	case pb.Code_E_PLAYER_FIRE:
		pf := &pb.EPlayerFire{}
		err = proto.Unmarshal(payload.Body, pf)
		if err == nil {
			log.Printf("%s EPlayerFire: %#v\n", d.user.NickName, pf)
		}
		matchId := d.match.GetMatchId()
		var damage int32 = 20
		otherGuyId := d.getOtherGuy()
		hit := &pb.CPlayerHit{}
		hit.MatchId = &matchId
		hit.P1 = &otherGuyId
		hit.P2 = &d.user.Id
		hit.Damage = &damage
		d.send(pb.Code_C_PLAYER_HIT, hit)
	case pb.Code_E_PLAYER_HIT:

	}
	if err != nil {
		log.Fatalf("%s OnRecievePayload Error: %s\n", d.user.NickName, err.Error())
	}

	return
}

func (d *DefaultClient) getOtherGuy() int64 {
	if d.match != nil {
		for _, v := range d.match.Players {
			if v.GetId() != d.user.Id {
				return v.GetId()
			}
		}
	}
	return 0
}

func (d *DefaultClient) handleErr(err error) {
	if err != nil {
		log.Fatalf("%s error: %s\n", d.user.NickName, err.Error())
	}
}
