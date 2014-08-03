package missle

import (
	"github.com/miraclew/mrs/mnet"
	"github.com/miraclew/mrs/pb"
	"log"
	"time"
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
	// Players   map[int64]*Player
	KeyPoints []*Point
	PlayersId []int64
	State     int
	TurnIdx   int
	turnTimer *time.Timer
	turns     int // turns aleady run

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
	clientsId := make([]int64, 0)
	for _, v := range playersId {
		clientId, ok := game.GetClientId(v)
		if ok {
			clientsId = append(clientsId, clientId)
		}
	}
	channelId, _ := manager.NewChannel(clientsId)

	turns := 10 * len(playersId)
	match := &Match{
		Id:        seq,
		ChannelId: channelId,
		PlayersId: playersId,
		State:     STATE_READY,
		manager:   manager,
		game:      game,
		turns:     turns,
	}
	match.InitKeyPoints()
	game.initMatchPlayers(match.Id, playersId)

	matchs[match.Id] = match
	return match, nil
}

func (m *Match) InitKeyPoints() {
	m.KeyPoints = MakeKeyPoints(16)
}

func GetMatch(id int64) *Match {
	match, ok := matchs[id]
	if !ok {
		log.Printf("Match: %d not found", id)
	}
	return match
}

func (m *Match) Begin() (err error) {
	if m.State != STATE_READY {
		err = NewMissleErr(ERR_INVALID_STATE, string(m.State))
		return
	}

	mi := &pb.EMatcInit{}
	mi.MatchId = &m.Id
	pbPlayers := make([]*pb.Player, 0)
	for _, p := range m.PlayersId {
		v := m.game.GetPlayer(p)
		player := &pb.Player{}
		player.Id = &v.Id
		player.NickName = &v.NickName
		player.Avatar = &v.Avatar
		player.IsLeft = &v.IsLeft
		player.Position = &pb.Point{X: &v.Position.X, Y: &v.Position.Y}
		player.Health = &v.Health
		pbPlayers = append(pbPlayers, player)
	}

	mi.Players = pbPlayers
	pbKeyPoints := make([]*pb.Point, 0)
	for _, v := range m.KeyPoints {
		pbKeyPoints = append(pbKeyPoints, &pb.Point{X: &v.X, Y: &v.Y})
	}
	mi.Points = pbKeyPoints

	msg := &mnet.Message{Code: pb.Code_E_MATCH_INIT, MSG: mi}

	err = m.manager.PushToChannel(m.ChannelId, msg)
	if err != nil {
		return
	}
	m.State = STATE_PLAYING
	m.NextTurn()
	return
}

func (m *Match) NextTurn() {
	if m.State == STATE_END {
		return
	}
	if m.turnTimer != nil {
		m.turnTimer.Stop()
	}
	// if m.turns <= 0 {
	// 	m.End()
	// }

	m.TurnIdx++
	if m.TurnIdx >= len(m.PlayersId) {
		m.TurnIdx = 0
	}

	playerId := m.PlayersId[m.TurnIdx]
	mt := &pb.EMatchTurn{}
	mt.MatchId = &m.Id
	mt.PlayerId = &playerId
	msg := &mnet.Message{Code: pb.Code_E_MATCH_TURN, MSG: mt}
	m.manager.PushToChannel(m.ChannelId, msg)

	// schedule next turn
	m.turnTimer = time.AfterFunc(time.Duration(5)*time.Second, m.NextTurn)
	m.turns--
}

func (m *Match) End() {
	log.Printf("MatchEnd: %#v\n", m.PlayersId)
	if m.State == STATE_END {
		return
	}
	m.State = STATE_END // should before send message, to avoid send twice

	for _, v := range m.PlayersId {
		player := m.game.GetPlayer(v)
		var point int32
		if player.Health <= 0 {
			point = -100
		} else {
			point = 100
		}
		player.MatchId = 0

		UpdatePlayerPoints(v, point)

		me := &pb.EMatchEnd{}
		me.MatchId = &v
		me.Points = &point
		msg := &mnet.Message{Code: pb.Code_E_MATCH_END, MSG: me}
		m.pushToUser(v, msg)
	}
}

func (m *Match) PlayerExit(playerId int64) error {
	log.Printf("PlayerExit: %d", playerId)
	if m.State == STATE_END {
		return NewMissleErr(ERR_INVALID_STATE, m.State)
	}

	m.End()
	return nil
}

func (m *Match) PlayerMove(playerId int64, pos Point) error {
	if m.State == STATE_END {
		return NewMissleErr(ERR_INVALID_STATE, m.State)
	}

	if !CheckPosition(pos) {
		return NewMissleErr(ERR_INVALID_POSITION, pos.X, pos.Y)
	}

	player := m.game.GetPlayer(playerId)
	player.Position = &pos

	pm := &pb.EPlayerMove{}
	pm.MatchId = &m.Id
	pm.PlayerId = &playerId
	pm.Position = &pb.Point{X: &pos.X, Y: &pos.Y}
	msg := &mnet.Message{Code: pb.Code_E_PLAYER_MOVE, MSG: pm}
	m.manager.PushToChannel(m.ChannelId, msg)

	return nil
}

func (m *Match) PlayerFire(playerId int64, pos Point, velocity Point) error {
	if m.State == STATE_END {
		return NewMissleErr(ERR_INVALID_STATE, m.State)
	}
	if playerId2 := m.PlayersId[m.TurnIdx]; playerId != playerId2 {
		log.Printf("PlayerFire, it's not your(%d) turn to fire, turn: %d", playerId2, playerId)
		return NewMissleErr(ERR_INVALID_STATE, m.State)
	}

	pf := &pb.EPlayerFire{}
	pf.MatchId = &m.Id
	pf.PlayerId = &playerId
	pf.Velocity = &pb.Point{X: &velocity.X, Y: &velocity.Y}
	msg := &mnet.Message{Code: pb.Code_E_PLAYER_FIRE, MSG: pf}
	m.manager.PushToChannel(m.ChannelId, msg)

	// wait a few time(missle hit) to trigger next turn
	return nil
}

// p1 hit p2
func (m *Match) PlayerHit(p1 int64, p2 int64, damage int32) error {
	if m.State == STATE_END {
		return NewMissleErr(ERR_INVALID_STATE, m.State)
	}
	// if playerId := m.PlayersId[m.TurnIdx]; p1 != playerId {
	// 	log.Printf("PlayerHit, it's not your(%d) turn to fire, turn: %d", p1, playerId)
	// 	return NewMissleErr(ERR_INVALID_STATE, m.State)
	// }

	if p2 > 0 && damage > 0 {
		newHealth, oldHealth := m.changeHealth(p2, -damage)
		player1 := m.game.GetPlayer(p1)
		player1.PointsWin += newHealth - oldHealth

		ph := &pb.EPlayerHit{}
		ph.MatchId = &m.Id
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

	// m.NextTurn()

	return nil
}

func (m *Match) shouldGameOver() bool {
	return true
}

func (m *Match) changeHealth(playerId int64, healthChange int32) (nh, oh int32) {
	player := m.game.GetPlayer(playerId)
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
		log.Printf("game.GetClientId(%d) failed, not online?\n", userId)
	}
}
