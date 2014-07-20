package missle

import (
	"encoding/json"
	"fmt"
)

// Interfaces this module depends on

type PushHandler interface {
	NewChannel(subsId []int64) (channelId int64, err error)
	PushToUser(userId int64, message interface{}) (err error)
	PushToChannel(chanelId int64, message interface{}) (err error)
}

type PusherMock struct{}

func (p *PusherMock) NewChannel(subsId []int64) (channelId int64, err error) {
	return 1, nil
}

func (p *PusherMock) PushToUser(userId int64, message interface{}) (err error) {
	bytes, _ := json.MarshalIndent(message, "", "  ")
	fmt.Printf("PushToUser: %d, %s\n", userId, string(bytes))
	return nil
}

func (p *PusherMock) PushToChannel(chanelId int64, message interface{}) (err error) {
	bytes, _ := json.MarshalIndent(message, "", "  ")
	fmt.Printf("PushToChannel: %d, %s\n", chanelId, string(bytes))
	return nil
}
