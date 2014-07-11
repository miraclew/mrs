package push

import (
// "fmt"
)

type Pusher struct {
}

// Push to one user
func PushToUser(userId int64, message *Message) (err error) {
	err = nil
	return
}

// Push to channel
func PushToChannel(chanelId int64, message *Message) (err error) {
	channel := GetChannel(chanelId)
	if channel != nil {
		for _, v := range channel.Subs {
			PushToUser(v, message)
		}
	}

	err = nil
	return
}
