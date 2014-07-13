package missle

import (
	"encoding/json"
	"fmt"
)

type Pusher interface {
	NewChannel(subsId []int64) (channelId int64, err error)
	// Push to one user
	PushToUser(userId int64, message interface{}) (err error)
	// Push to channel
	PushToChannel(chanelId int64, message interface{}) (err error)

	// IsConnected(userId int64)
	ConnectionHandle(handler PushHandler)
}

type PushHandler interface {
	OnValidateToken(token string) int64 // token to userId
	OnConnected(userId int64)
	OnDisconnected(userId int64)
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

func (p *PusherMock) ConnectionHandle(handler PushHandler) {

}
