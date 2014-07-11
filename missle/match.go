package missle

import (
// "github.com/miraclew/mrs/push"
)

const (
	STATE_READY   = 1
	STATE_PLAYING = 2
	STATE_END     = 3
)

type Match struct {
	Id        int64
	ChannelId int64
	Players   []*Player
	State     int
	TurnIdx   int
}

var seq int64 = 0
var matchs map[int64]*Match

func NewMatch(playersId []int64) *Match {
	seq++
	match := &Match{
		Id:        seq,
		ChannelId: GetChannelId(),
		Players:   makePlayers(playersId),
		State:     STATE_READY,
	}

	matchs[match.Id] = match
	return match
}

func makePlayers(playersId []int64) []*Player {
	players := make([]*Player, len(playersId))
	isLeft := true
	for i := 0; i < len(playersId); i++ {
		playerId := playersId[i]
		profile := GetMatch(playerId)
		pos := MakePositionFor(isLeft, 0)
		players[i] = &Player{playerId, profile.Nickname, profile.Avatar, isLeft, pos}

		isLeft = !isLeft
	}
	return players
}

func GetMatch(id int64) *Match {
	return matchs[id]
}

func (m *Match) Begin() (err error) {
	if m.State != STATE_READY {
		err = NewMissleErr(ERR_INVALID_STATE, string(m.State))
		return
	}

	keyPoints := MakeKeyPoints(16)
	msg := &Message{}
	msg.Header.Name = MN_MatchBegin
	msg.Header.ChannelId = m.ChannelId
	msg.Body = &MatchBegin{Players: m.Players, KeyPoints: keyPoints}

	PushToChannel(m.ChannelId, msg)
}

func (m *Match) NextTurn() {
	m.TurnIdx++
	if m.TurnIdx >= len(m.Players) {
		m.TurnIdx = 0
	}

	playerId = m.Players[m.TurnIdx].Id

	msg := &Message{}
	msg.Header.Name = MN_MatchTurn

	PushToUser(playerId, msg)
}

func (m *Match) End() {

}

func (m *Match) PlayerMove(playerId int64, pos Point) {
	msg := &Message{}
	msg.Header.Name = MN_PlayerMove
	msg.Header.ChannelId = m.ChannelId
	msg.Body = &PlayerMove{playerId, pos}

	PushToChannel(m.ChannelId, msg)
}

func (m *Match) PlayerFire() {

}

func (m *Match) PlayerHealth(playerId int64, healthChange int) {
	newHealth := m.changeHealth(playerId, healthChange)
	if newHealth == 0 {

	}

	msg := &Message{}
	msg.Header.Name = MN_PlayerHealth
	msg.Header.ChannelId = m.ChannelId
	msg.Body = &PlayerHealth{playerId, healthChange}

	PushToChannel(m.ChannelId, msg)
}

func (m *Match) shouldGameOver() {

}

func (m *Match) changeHealth(playerId int64, healthChange int) int {
	for _, v := range m.Players {
		if v.Id == playerId {
			v.Health += healthChange
			if v.Health < 0 {
				v.Health = 0
			}

			return v.Health
		}
	}

	return -1
}
