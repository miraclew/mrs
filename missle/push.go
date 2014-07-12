package missle

import (
	"encoding/json"
	"fmt"
)

var channelId int64 = 1

func PushToUser(userId int64, m *Message) {
	bytes, _ := json.Marshal(m)
	fmt.Printf("PushToUser: %d, %s", userId, string(bytes))
}

func PushToChannel(channelId int64, m *Message) {
	bytes, _ := json.Marshal(m)
	fmt.Printf("PushToChannel: %d, %s", channelId, string(bytes))
}

func GetChannelId() int64 {
	channelId++
	return channelId
}
