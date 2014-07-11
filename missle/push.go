package missle

import (
	"encoding/json"
	"fmt"
)

var chanelId int64 = 1

func PushToUser(userId int64, m *Message) {
	_, bytes := json.Marshal(m)
	fmt.Printf("PushToUser: %d, $s", userId, string(bytes))
}

func PushToChannel(chanelId int64, m *Message) {
	_, bytes := json.Marshal(m)
	fmt.Printf("PushToChannel: %d, $s", channelId, string(bytes))
}

func GetChannelId() int64 {
	chanelId++
	return chanelId
}
