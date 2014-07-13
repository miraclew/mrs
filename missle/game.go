package missle

type Game struct {
	waitQueue []int64           // waiting players
	players   map[int64]*Player // all online players
}

func (g *Game) Init() {
	g.waitQueue = []int64{}
	g.players = make(map[int64]*Player)
}

func (g *Game) PlayerEnter(playerId int64) {
	if len(g.waitQueue) > 0 {
		p1 := g.waitQueue[0]
		p2 := playerId
		g.waitQueue = g.waitQueue[1:]

		NewMatch([]int64{p1, p2}, nil)
	} else {
		g.waitQueue = append(g.waitQueue, playerId)
	}
}

func (g *Game) PlayerExit(playerId int64) {
	delete(g.players, playerId)
}

func (g *Game) initPlayer(playerId int64) (player *Player) {
	profile := FindUserById(playerId)
	player = &Player{Id: playerId, NickName: profile.UserName, Avatar: profile.Avatar}
	g.players[playerId] = player
	return
}
