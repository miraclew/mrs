package missle

import (
	"github.com/miraclew/mrs/mnet"
	"github.com/miraclew/mrs/pb"
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

	manager *mnet.Manager
	game    *Game
}

var seq int64 = 0
var matchs = make(map[int64]*Match)

func NewMatch(game *Game, playersId []int64, manager *mnet.Manager) (*Match, error) {
	if playersId == nil || len(playersId) < 2 {
		return nil, NewMissleErr(ERR_INVALID_ARGS, "playersId is nil or less than 2")
	}

	if manager == nil {
		return nil, NewMissleErr(ERR_INVALID_ARGS, "manager is nil")
	}

	log.Printf("NewMatch(%#v, %#v)", playersId, manager)
	seq++
	channelId, _ := manager.NewChannel(playersId)
	match := &Match{
		Id:        seq,
		ChannelId: channelId,
		Players:   makePlayers(playersId),
		PlayersId: playersId,
		State:     STATE_READY,
		manager:   manager,
		game:      game,
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

	// keyPoints := MakeKeyPoints(16)
	var players []*Player
	for _, v := range m.Players {
		players = append(players, v)
	}

	mi := &pb.EMatcInit{}
	mi.MatchId = &m.Id
	pbPlayers := make([]*pb.Player, len(m.Players))
	mi.Players = pbPlayers

	msg := &mnet.Message{Code: pb.Code_E_MATCH_INIT, MSG: mi}

	m.manager.PushToChannel(m.ChannelId, msg)
	m.State = STATE_PLAYING
	return nil
}

func (m *Match) NextTurn() {
	m.TurnIdx++
	if m.TurnIdx >= len(m.Players) {
		m.TurnIdx = 0
	}

	playerId := m.PlayersId[m.TurnIdx]
	mt := &pb.EMatchTurn{}
	mt.PlayerId = &playerId
	msg := &mnet.Message{Code: pb.Code_E_MATCH_TURN, MSG: mt}
	m.pushToUser(playerId, msg)
}

func (m *Match) End() {
	for _, v := range m.Players {
		var point int32
		if v.Health == 0 {
			point = -100
		} else {
			point = 100
		}

		UpdatePlayerPoints(v.Id, point)

		me := &pb.EMatchEnd{}
		me.Points = &point
		msg := &mnet.Message{Code: pb.Code_E_MATCH_END, MSG: me}
		m.pushToUser(v.Id, msg)
	}
	m.State = STATE_END
}

func (m *Match) PlayerMove(playerId int64, pos Point) error {
	if !CheckPosition(pos) {
		return NewMissleErr(ERR_INVALID_POSITION, pos.X, pos.Y)
	}

	player := m.Players[playerId]
	player.Position = pos

	pm := &pb.EPlayerMove{}
	pm.PlayerId = &playerId
	pm.Position = &pb.Point{X: &pos.X, Y: &pos.Y}
	msg := &mnet.Message{Code: pb.Code_E_PLAYER_MOVE, MSG: pm}
	m.manager.PushToChannel(m.ChannelId, msg)

	return nil
}

func (m *Match) PlayerFire(playerId int64, pos Point, velocity Point) {
	pf := &pb.EPlayerFire{}
	pf.PlayerId = &playerId
	pf.Velocity = &pb.Point{X: &velocity.X, Y: &velocity.Y}
	msg := &mnet.Message{Code: pb.Code_E_PLAYER_FIRE, MSG: pf}
	m.manager.PushToChannel(m.ChannelId, msg)
}

// p1 hit p2
func (m *Match) PlayerHit(p1 int64, p2 int64, damage int32) {
	newHealth, oldHealth := m.changeHealth(p2, -damage)
	player1 := m.Players[p1]
	player1.PointsWin += newHealth - oldHealth

	ph := &pb.EPlayerHit{}
	ph.P1 = &p1
	ph.P2 = &p2
	ph.Damage = &damage
	msg := &mnet.Message{Code: pb.Code_E_PLAYER_HIT, MSG: ph}

	m.manager.PushToChannel(m.ChannelId, msg)

	if newHealth == 0 {
		if m.shouldGameOver() {
			m.End()
		}
	}
}

func (m *Match) shouldGameOver() bool {
	return true
}

func (m *Match) changeHealth(playerId int64, healthChange int32) (nh, oh int32) {
	player := m.Players[playerId]
	oh = player.Health

	player.Health += healthChange
	if player.Health < 0 {
		player.Health = 0
	}

	nh = player.Health
	return
}

func (m *Match) pushToUser(userId int64, msg *mnet.Message) {
	clientId, ok := m.game.GetClientId(userId)
	if ok {
		m.manager.PushToClient(clientId, msg)
	} else {
		log.Println("pushToUser failed, not online?")
	}
}
