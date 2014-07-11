package push

import (
// "gopkg.in/fatih/set.v0" // https://github.com/fatih/set
)

type Channel struct {
	Id   int64
	Subs []int64
}

var seq int64 = 0
var channels map[int64]*Channel

func NewChannel(subs []int64) *Channel {
	seq++
	channel := &Channel{
		Id:   seq,
		Subs: subs,
	}

	channels[channel.Id] = channel
	return channel
}

func GetChannel(id int64) *Channel {
	return channels[id]
}
