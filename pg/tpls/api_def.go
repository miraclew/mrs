package tpls

type Player struct {
	UserId int64
	Name   string
	Age    int
	Points int
}

type PlayerService interface {
	GetPlayer(userId int64) (player *Player)
	GetOnlinePlayers() (players []*Player)
}
