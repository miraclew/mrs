package push

import (
	"gopkg.in/fatih/set.v0" // https://github.com/fatih/set
)

type PushChannel struct {
	Id   int64
	Subs set.Set
}

var seq int64 = 0
var channels map[int64]*PushChannel

func NewChannel(subs set.Set) *PushChannel {
	seq++
	channel := &PushChannel{
		Id:   seq,
		Subs: subs,
	}

	channels[channel.Id] = channel
}

func GetChannel(id int64) *PushChannel {
	return channels[id]
}
