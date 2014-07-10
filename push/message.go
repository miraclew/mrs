package push

const (
	NameMatchBegin   = "MatchBegin"
	NameMatchEnd     = "MatchEnd"
	NameMatchTurn    = "MatchTurn"
	NamePlayerMove   = "PlayerMove"
	NamePlayerFire   = "PlayerFire"
	NamePlayerHealth = "PlayerHealth"
)

type Header struct {
	Id        int64
	Name      string
	MatchId   int64
	ChannelId int64
}

type Message struct {
	Header Header
	Body   interface{}
}

type Point struct {
	X float32
	Y float32
}

/* Message body definition */

// MatchBegin, send to all players
type MatchBegin struct {
	Players []struct {
		Id       int64
		NickName string
		Avatar   string
		IsLeft   bool
		Position Point
	}

	KeyPoints []Point
}

// Send to each player
type MatchEnd struct {
	Point Point
}

type MatchTurn struct{}

type PlayerMove struct {
	Position Point
}

type PlayerFire struct {
	Velocity Point
}

type PlayerHealth struct {
	Health int
}
