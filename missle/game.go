package missle

type Game struct {
	waitQueue []int64
}

func (g *Game) Init() {
	g.waitQueue = make([]int64)
}

func (g *Game) Enter(playerId int64) {
	append(g.waitQueue, playerId)
}
