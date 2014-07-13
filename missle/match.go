package missle

import (
	"log"
)

const (
	STATE_READY   = 1
	STATE_PLAYING = 2
	STATE_END     = 3
)

type MatchPlayer struct {
	Player
}

type Match struct {
	Id        int64
	ChannelId int64
	Players   map[int64]*Player
	PlayersId []int64
	State     int
	TurnIdx   int
	Pusher    Pusher
}

var seq int64 = 0
var matchs = make(map[int64]*Match)

func NewMatch(playersId []int64, pusher Pusher) (*Match, error) {
	if playersId == nil || len(playersId) < 2 {
		return nil, NewMissleErr(ERR_INVALID_ARGS, "playersId is nil or less than 2")
	}

	if pusher == nil {
		return nil, NewMissleErr(ERR_INVALID_ARGS, "pusher is nil")
	}

	log.Printf("NewMatch(%#v, %#v)", playersId, pusher)
	seq++
	channelId, _ := pusher.NewChannel(playersId)
	match := &Match{
		Id:        seq,
		ChannelId: channelId,
		Players:   makePlayers(playersId),
		PlayersId: playersId,
		State:     STATE_READY,
		Pusher:    pusher,
	}

	matchs[match.Id] = match
	return match, nil
}

func makePlayers(playersId []int64) map[int64]*Player {
	players := make(map[int64]*Player)
	isLeft := true
	for i := 0; i < len(playersId); i++ {
		playerId := playersId[i]
		profile := FindUserById(playerId)
		pos := MakePositionFor(isLeft, 0)
		players[playerId] = &Player{playerId, profile.UserName, profile.Avatar, isLeft, *pos, 100, 0}

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
	m.Pusher.PushToChannel(m.ChannelId, msg)

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
	m.Pusher.PushToUser(playerId, msg)
}

func (m *Match) End() {
	for _, v := range m.Players {
		var point int
		if v.Health == 0 {
			point = -100
		} else {
			point = 100
		}

		UpdatePlayerPoints(v.Id, point)

		msg := m.newMessage(MN_MatchEnd, &MatchEnd{point})
		m.Pusher.PushToUser(v.Id, msg)
	}
	m.State = STATE_END
}

func (m *Match) PlayerMove(playerId int64, pos Point) error {
	if !CheckPosition(pos) {
		return NewMissleErr(ERR_INVALID_POSITION, pos.X, pos.Y)
	}

	player := m.Players[playerId]
	player.Position = pos

	msg := m.newMessage(MN_PlayerMove, &PlayerMove{playerId, pos})
	m.Pusher.PushToChannel(m.ChannelId, msg)

	return nil
}

func (m *Match) PlayerFire(playerId int64, pos Point, velocity Point) {
	msg := m.newMessage(MN_PlayerFire, &PlayerFire{playerId, velocity})
	m.Pusher.PushToChannel(m.ChannelId, msg)
}

// p1 hit p2
func (m *Match) PlayerAttack(p1 int64, p2 int64, damage int) {
	newHealth, oldHealth := m.changeHealth(p2, -damage)
	player1 := m.Players[p1]
	player1.PointsWin += newHealth - oldHealth

	msg := m.newMessage(MN_PlayerHealth, &PlayerHealth{p2, newHealth})
	m.Pusher.PushToChannel(m.ChannelId, msg)

	if newHealth == 0 {
		if m.shouldGameOver() {
			m.End()
		}
	}
}

func (m *Match) shouldGameOver() bool {
	return true
}

func (m *Match) changeHealth(playerId int64, healthChange int) (nh, oh int) {
	player := m.Players[playerId]
	oh = player.Health

	player.Health += healthChange
	if player.Health < 0 {
		player.Health = 0
	}

	nh = player.Health
	return
}

func (m *Match) newMessage(name string, body interface{}) *Message {
	msg := &Message{}
	msg.Header.Name = name
	msg.Header.ChannelId = m.ChannelId
	msg.Header.MatchId = m.Id
	msg.Body = body
	return msg
}
