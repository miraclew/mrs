package push

import (
// "fmt"
)

type Pusher struct {
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
