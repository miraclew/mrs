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
	waitQueue *WaitQueue        // waiting players
	players   map[int64]*Player // all online players
	c2uMap    map[int64]int64   // map client.Id => userId
	u2cMap    map[int64]int64   // map userId => client.Id
	manager   *mnet.Manager
	Db        *gorp.DbMap

	seqUserId int64
}

var game *Game

func init() {
}

func NewGame(manager *mnet.Manager) *Game {
	game = &Game{}
	game.init()
	game.manager = manager
	game.Db = model.InitDb(DSN)
	game.seqUserId = 1
	manager.Handler = game
	return game
}

func (g *Game) init() {
	g.waitQueue = &WaitQueue{}
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
	userId, ok := g.c2uMap[clientId]
	if ok {
		delete(g.c2uMap, clientId)
		delete(g.u2cMap, userId)
		g.waitQueue.Delete(userId)
		player := g.GetPlayer(userId)
		log.Printf("player disconnected: %#v", player)
		if player != nil && player.MatchId != 0 {
			match := GetMatch(player.MatchId)
			if match != nil {
				match.PlayerExit(userId)
			}
			player.MatchId = 0
		}
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
		user := model.User{}
		// err = g.Db.SelectOne(&user, "select id from users where UserName=? and Password=?", auth.GetUserName(), auth.GetPassword())
		err = g.Db.SelectOne(&user, "select * from users where Id=?", g.seqUserId)
		g.seqUserId++
		if g.seqUserId > 9 {
			g.seqUserId = 1
		}

		var code int32 = 0
		if err != nil {
			code = ERR_INVALID_CREDENTIAL
			log.Printf("Auth failed: (username=%s)", auth.GetUserName())
		} else {
			log.Printf("Auth success: (username=%s) clientId:%d userId:%d", user.UserName, clientId, user.Id)
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
	} else if code == pb.Code_C_MATCH_EXIT {
		exit := &pb.CMatchExit{}
		err = proto.Unmarshal(payload.Body, exit)
		if err == nil {
			if match := GetMatch(exit.GetMatchId()); match != nil {
				match.PlayerExit(playerId)
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
	if g.waitQueue.Len() > 0 {
		p1, _ := g.waitQueue.Pop().(int64)
		p2 := playerId
		if p1 == p2 {
			g.waitQueue.Push(playerId)
			err = NewMissleErr(ERR_INVALID_STATE, fmt.Sprintf("userId: %d already enter game", p1))
			return
		}

		log.Printf("NewMatch for p1=%d, p2=%d", p1, p2)
		var match *Match
		match, err = NewMatch(g, []int64{p1, p2}, g.manager)
		if err != nil {
			log.Printf("NewMatch failed: %s", err.Error())
			return
		}
		match.Begin()
	} else {
		g.waitQueue.Push(playerId)
	}

	log.Printf("WaitQueue %s\n", g.waitQueue)

	return
}

func (g *Game) PlayerExit(playerId int64) {

}

func (g *Game) GetPlayer(playerId int64) *Player {
	player, ok := g.players[playerId]
	if ok {
		return player
	}
	return nil
}

func (g *Game) initMatchPlayers(matchId int64, playersId []int64) {
	isLeft := true
	for i := 0; i < len(playersId); i++ {
		playerId := playersId[i]
		player := g.GetPlayer(playerId)
		if player == nil {
			player, _ = g.initPlayer(playerId)
		}
		player.IsLeft = isLeft
		player.MatchId = matchId
		player.Health = 100
		player.Position = MakePositionFor(isLeft, 0)
		g.players[playerId] = player
		isLeft = !isLeft
	}
}

func (g *Game) initPlayer(playerId int64) (player *Player, err error) {
	user := model.User{}
	err = g.Db.SelectOne(&user, "select * from users where Id=?", playerId)
	if err != nil {
		log.Fatalf("Find user err: %s", err.Error())
		return
	}

	player = &Player{playerId, user.UserName, user.Avatar, 0, true, &Point{}, 100, 0}
	g.players[playerId] = player
	return
}
