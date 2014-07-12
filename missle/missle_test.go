package missle

import (
	"encoding/json"
	"fmt"
	"testing"
)

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

func TestA(t *testing.T) {
	t.Skip("...")
	// var x float32
	// x = 1.0 * 2
	// y := x * 2
	// fmt.Printf("%f, %f", x, y)
	points := MakeKeyPoints(16)
	for _, p := range points {
		fmt.Printf("%f, %f\n", p.X, p.Y)
	}
}

func TestB(t *testing.T) {
	// t.Skip("...")
	var pusher Pusher
	pusher = &PusherMock{}
	match := NewMatch([]int64{1, 2}, pusher)
	match.Begin()

	if match.State != STATE_PLAYING {
		t.Error("STATE error")
	}

	match.PlayerMove(1, Point{0.44, 0.34})
	match.PlayerFire(2, Point{0.5, 0.6}, Point{0.1, 0.2})
	match.PlayerHealth(1, -10)
	match.PlayerHealth(2, +10)

	match.PlayerHealth(1, -90)
}
