package missle

const (
	MN_MatchBegin   = "MatchBegin"
	MN_MatchEnd     = "MatchEnd"
	MN_MatchTurn    = "MatchTurn"
	MN_PlayerMove   = "PlayerMove"
	MN_PlayerFire   = "PlayerFire"
	MN_PlayerHealth = "PlayerHealth"
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

/* Message body definition */

// MatchBegin, send to all players
type MatchBegin struct {
	Players []*Player

	KeyPoints []Point
}

// Send to each player
type MatchEnd struct {
	Point Point
}

type MatchTurn struct {
	PlayerId int64
}

type PlayerMove struct {
	PlayerId int64
	Position Point
}

type PlayerFire struct {
	Velocity Point
}

type PlayerHealth struct {
	PlayerId int64
	Health   int
}

func SendMessage(msg *Message) {
	fmt.Printf("Send: ")
}
