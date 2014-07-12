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
	Players   map[int64]*Player
	PlayersId []int64
	State     int
	TurnIdx   int
}

var seq int64 = 0
var matchs = make(map[int64]*Match)

func NewMatch(playersId []int64) *Match {
	seq++
	match := &Match{
		Id:        seq,
		ChannelId: GetChannelId(),
		Players:   makePlayers(playersId),
		PlayersId: playersId,
		State:     STATE_READY,
	}

	matchs[match.Id] = match
	return match
}

func makePlayers(playersId []int64) map[int64]*Player {
	players := make(map[int64]*Player)
	isLeft := true
	for i := 0; i < len(playersId); i++ {
		playerId := playersId[i]
		profile := GetProfile(playerId)
		pos := MakePositionFor(isLeft, 0)
		players[playerId] = &Player{playerId, profile.Nickname, profile.Avatar, isLeft, *pos, 100}

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
	var players []*Player
	for _, v := range m.Players {
		players = append(players, v)
	}

	msg := m.newMessage(MN_MatchBegin, &MatchBegin{players, keyPoints})
	PushToChannel(m.ChannelId, msg)

	m.State = STATE_PLAYING
	return nil
}

func (m *Match) NextTurn() {
	m.TurnIdx++
	if m.TurnIdx >= len(m.Players) {
		m.TurnIdx = 0
	}

	playerId := m.PlayersId[m.TurnIdx]

	msg := m.newMessage(MN_MatchTurn, nil)
	PushToUser(playerId, msg)
}

func (m *Match) End() {
	for _, v := range m.Players {
		var point int
		if v.Health == 0 {
			point = -100
		} else {
			point = 100
		}

		msg := m.newMessage(MN_MatchEnd, &MatchEnd{point})
		PushToUser(v.Id, msg)
	}
	m.State = STATE_END
}

func (m *Match) PlayerMove(playerId int64, pos Point) {
	msg := m.newMessage(MN_PlayerMove, &PlayerMove{playerId, pos})
	PushToChannel(m.ChannelId, msg)
}

func (m *Match) PlayerFire(playerId int64, pos Point, velocity Point) {
	msg := m.newMessage(MN_PlayerFire, &PlayerFire{playerId, velocity})
	PushToChannel(m.ChannelId, msg)
}

func (m *Match) PlayerHealth(playerId int64, healthChange int) {
	msg := m.newMessage(MN_PlayerHealth, &PlayerHealth{playerId, healthChange})
	PushToChannel(m.ChannelId, msg)

	newHealth := m.changeHealth(playerId, healthChange)
	if newHealth == 0 {
		if m.shouldGameOver() {
			m.End()
		}
	}
}

func (m *Match) shouldGameOver() bool {
	return true
}

func (m *Match) changeHealth(playerId int64, healthChange int) int {
	player := m.Players[playerId]
	if player == nil {
		return -1
	}
	player.Health += healthChange
	if player.Health < 0 {
		player.Health = 0
	}

	return player.Health
}

func (m *Match) newMessage(name string, body interface{}) *Message {
	msg := &Message{}
	msg.Header.Name = MN_PlayerHealth
	msg.Header.ChannelId = m.ChannelId
	msg.Header.MatchId = m.Id
	msg.Body = body
	return msg
}
