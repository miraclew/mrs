package missle

import (
	"fmt"
	"log"
)

type Game struct {
	waitQueue []int64           // waiting players
	players   map[int64]*Player // all online players
	pusher    Pusher
}

var game *Game

func init() {
	game = &Game{}
	game.init()
}

func GetGame() *Game {
	return game
}

func (g *Game) init() {
	g.waitQueue = []int64{}
	g.players = make(map[int64]*Player)
}

func (g *Game) SetPuser(pusher Pusher) {
	g.pusher = pusher
	pusher.ConnectionHandle(g)
}

/* ConnectionHandler */
func (g *Game) OnValidateToken(token string) int64 {
	uid, err := GetUidByToken(token)
	if err != nil {
		uid = 0
	}

	return uid
}

func (g *Game) OnConnected(playerId int64) {
	player, _ := g.initPlayer(playerId)
	g.players[playerId] = player
}

func (g *Game) OnDisconnected(playerId int64) {
	delete(g.players, playerId)
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
		match, err = NewMatch([]int64{p1, p2}, g.pusher)
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
