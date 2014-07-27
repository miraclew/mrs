package missle

import (
	"code.google.com/p/goprotobuf/proto"
	"fmt"
	"github.com/coopernurse/gorp"
	"github.com/miraclew/mrs/missle/model"
	"github.com/miraclew/mrs/mnet"
	"github.com/miraclew/mrs/pb"
	"log"
)

type Game struct {
	waitQueue []int64           // waiting players
	players   map[int64]*Player // all online players
	c2uMap    map[int64]int64   // map client.Id => userId
	u2cMap    map[int64]int64   // map userId => client.Id
	manager   *mnet.Manager
	Db        *gorp.DbMap
}

var game *Game

func init() {
}

func NewGame(manager *mnet.Manager) *Game {
	game = &Game{}
	game.init()
	game.manager = manager
	game.Db = model.InitDb(DSN)
	manager.Handler = game
	return game
}

func (g *Game) init() {
	g.waitQueue = []int64{}
	g.players = make(map[int64]*Player)
	g.c2uMap = make(map[int64]int64)
	g.u2cMap = make(map[int64]int64)
}

func (g *Game) GetClientId(userId int64) (clientId int64, ok bool) {
	clientId, ok = g.u2cMap[userId]
	return
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
	// player, _ := g.initPlayer(playerId)
	// g.players[playerId] = player
}

func (g *Game) OnDisconnected(clientId int64) {
	// delete(g.players, playerId)
	userId, ok := g.c2uMap[clientId]
	if ok {
		delete(g.c2uMap, clientId)
		delete(g.u2cMap, userId)
	}
	log.Printf("c2uMap= %#v\n", g.c2uMap)
	log.Printf("u2cMap= %#v\n", g.u2cMap)
}

func (g *Game) OnRecievePayload(clientId int64, payload *mnet.Payload) {
	var err error
	playerId := g.c2uMap[clientId]
	code := pb.Code(payload.Code)
	log.Printf("Client(%d) playersId(%d) recv payload: %#v", clientId, playerId, code.String())

	if code == pb.Code_C_AUTH {
		auth := &pb.CAuth{}
		err = proto.Unmarshal(payload.Body, auth)
		if err == nil {
			log.Printf("CAuth: %s", auth.String())
		}
		eauth := &pb.EAuth{}
		//user := FindUserByCredential(auth.GetUserName(), auth.GetPassword())
		user := model.User{}
		err = g.Db.SelectOne(&user, "select id from users where UserName=? and Password=?", auth.GetUserName(), auth.GetPassword())
		var code int32 = 0
		if err != nil {
			code = ERR_INVALID_CREDENTIAL
			log.Printf("Auth failed: (username=%s)", auth.GetUserName())
		} else {
			log.Printf("Auth success: (username=%s) clientId:%d userId:%d", auth.GetUserName(), clientId, user.Id)
			g.c2uMap[clientId] = user.Id
			g.u2cMap[user.Id] = clientId
		}

		eauth.Code = &code
		eauth.UserId = &user.Id

		msg := &mnet.Message{Code: pb.Code_E_AUTH, MSG: eauth}
		g.manager.PushToClient(clientId, msg)
	} else if code == pb.Code_C_MATCH_ENTER {
		matchEnter := &pb.CMatchEnter{}
		err = proto.Unmarshal(payload.Body, matchEnter)
		g.PlayerEnter(playerId)
	} else if code == pb.Code_C_PLAYER_MOVE {
		move := &pb.CPlayerMove{}
		err = proto.Unmarshal(payload.Body, move)

		match := GetMatch(move.GetMatchId())
		if match != nil {
			match.PlayerMove(playerId, Point{X: move.GetPosition().GetX(), Y: move.GetPosition().GetY()})
		}
	} else if code == pb.Code_C_PLAYER_FIRE {
		fire := &pb.CPlayerFire{}
		err = proto.Unmarshal(payload.Body, fire)

		match := GetMatch(fire.GetMatchId())
		if match != nil {
			match.PlayerFire(playerId, Point{}, Point{X: fire.GetVelocity().GetX(), Y: fire.GetVelocity().GetY()})
		}
	} else if code == pb.Code_C_PLAYER_HIT {
		hit := &pb.CPlayerHit{}
		err = proto.Unmarshal(payload.Body, hit)
		if err == nil {
			log.Printf("%s", hit.String())
			match := GetMatch(hit.GetMatchId())
			if match != nil {
				match.PlayerHit(hit.GetP1(), hit.GetP2(), hit.GetDamage())
			}
		}
	} else {
		log.Printf("Error: unknown command %d", code)
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

		log.Printf("NewMatch for p1=%d, p2=%d", p1, p2)
		var match *Match
		match, err = NewMatch(g, []int64{p1, p2}, g.manager)
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
