package missle

import (
	"code.google.com/p/goprotobuf/proto"
	"fmt"
	"github.com/miraclew/mrs/mnet"
	"github.com/miraclew/mrs/pb"
	"log"
)

type Game struct {
	waitQueue []int64           // waiting players
	players   map[int64]*Player // all online players
	clients   map[int64]int64   // map userId <=> client.Id
	manager   *mnet.Manager
}

var game *Game

func init() {
}

func NewGame(manager *mnet.Manager) *Game {
	game = &Game{}
	game.init()
	game.nextClientId = 1
	game.manager = manager
	manager.Handler = game
	return game
}

func (g *Game) init() {
	g.waitQueue = []int64{}
	g.players = make(map[int64]*Player)
}

/* ConnectionHandler */
func (g *Game) OnValidateToken(token string) int64 {
	uid, err := GetUidByToken(token)
	if err != nil {
		uid = 0
	}

	return uid
}

func (g *Game) OnConnected(clientId int64) {
	player, _ := g.initPlayer(playerId)
	g.players[playerId] = player
}

func (g *Game) OnDisconnected(clientId int64) {
	delete(g.players, playerId)
}

func (g *Game) OnRecievePayload(clientId int64, payload *mnet.Payload) {
	var err error
	code := pb.Code(payload.Code)
	if code == pb.Code_C_AUTH {
		auth := &pb.CAuth{}
		err = proto.Unmarshal(payload.Body, auth)
		eauth := &pb.EAuth{}
		user := FindUserByCredential(auth.GetUserName(), auth.GetPassword())
		var code int32 = 0
		if user == nil {
			code = -1
		}
		eauth.Code = &code

		msg := &mnet.Message{Code: pb.Code_E_AUTH, MSG: eauth}
		g.manager.PushToUser(playerId, msg)
	} else if code == pb.Code_C_MATCH_ENTER {
		matchEnter := &pb.CMatchEnter{}
		err = proto.Unmarshal(payload.Body, matchEnter)
		g.PlayerEnter(playerId)
	} else if code == pb.Code_C_PLAYER_MOVE {
		move := &pb.CPlayerMove{}
		err = proto.Unmarshal(payload.Body, move)

		match := GetMatch(move.GetMatchId())
		match.PlayerMove(playerId, Point{X: move.GetPosition().GetX(), Y: move.GetPosition().GetY()})
	} else if code == pb.Code_C_PLAYER_FIRE {
		fire := &pb.CPlayerFire{}
		err = proto.Unmarshal(payload.Body, fire)

		match := GetMatch(fire.GetMatchId())
		match.PlayerFire(playerId, Point{}, Point{X: fire.GetVelocity().GetX(), Y: fire.GetVelocity().GetY()})
	} else if code == pb.Code_C_PLAYER_HIT {
		hit := &pb.CPlayerHit{}
		err = proto.Unmarshal(payload.Body, hit)
		match := GetMatch(hit.GetMatchId())
		match.PlayerHit(hit.GetP1(), hit.GetP2(), hit.GetDamage())
	} else if code == pb.Code_C_PLAYER_HEALTH {

	}

	if err != nil {
		log.Printf("OnRecievePayload error:%s", err.Error())
	}
}

// Player enter game (connected)
func (g *Game) PlayerEnter(playerId int64) (err error) {
	if len(g.waitQueue) > 0 {
		p1 := g.waitQueue[0]
		p2 := playerId
		if p1 == p2 {
			err = NewMissleErr(ERR_INVALID_STATE, fmt.Sprintf("userId: %d already enter game", p1))
			return
		}
		g.waitQueue = g.waitQueue[1:]

		var match *Match
		match, err = NewMatch([]int64{p1, p2}, g.manager)
		if err != nil {
			log.Printf("NewMatch failed: %s", err.Error())
			return
		}
		match.Begin()
	} else {
		g.waitQueue = append(g.waitQueue, playerId)
	}

	return
}

func (g *Game) PlayerExit(playerId int64) {

}

func (g *Game) initPlayer(playerId int64) (player *Player, err error) {
	profile := FindUserById(playerId)
	if profile == nil {
		err = NewMissleErr(ERR_DATA_NOT_FOUND, fmt.Sprintf("userId: %d", playerId))
		return
	}

	//log.Printf("initPlayer %d profile: %#v \n", playerId, profile)
	player = &Player{Id: playerId, NickName: profile.UserName, Avatar: profile.Avatar}
	return
}
