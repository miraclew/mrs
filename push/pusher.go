package push

type Pusher struct {
}

// Push to one user
func (p *Pusher) PushToUser(userId int64, message *Message) (err error) {

}

// Push to channel
func (p *Pusher) PushToChannel(chanelId int64, message *Message) (err error) {
	channel := GetChannel(chanelId)
	if channel {
		subs := channel.Subs.List()
		for e := subs.Front(); e != nil; e = e.Next() {
			PushToUser(e.Value, message)
		}
	}
}
