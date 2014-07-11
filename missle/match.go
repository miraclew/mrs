package missle

import (
// "github.com/miraclew/mrs/push"
)

const (
	STATE_CREATED = 0
	STATE_READY   = 1
	STATE_PLAYING = 2
	STATE_END     = 3
)

type Match struct {
	Id      int64
	Players []*Player
	State   int
}

var seq int64 = 0
var matchs map[int64]*Match

func NewMatch(subs []int64) *Match {
	seq++
	match := &Match{
		Id: seq,
	}

	matchs[match.Id] = match
	return match
}

func GetMatch(id int64) *Match {
	return matchs[id]
}

func (m *Match) Begin() {
	msg := Message{}
	msg.Header.Id = 1
	msg.Header.Name = MN_PlayerMove
	msg.Body = PlayerMove{Position: Point{X: 11, Y: 22}}

	// PushToUser(1, msg)
}

func (m *Match) Turn() {

}

func (m *Match) End() {

}

func (m *Match) PlayerMove() {

}

func (m *Match) PlayerFire() {

}

func (m *Match) PlayerHealth() {

}
