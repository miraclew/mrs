package missle

type Pusher interface {
	NewChannel(subsId []int64) (channelId int64, err error)
	// Push to one user
	PushToUser(userId int64, message interface{}) (err error)
	// Push to channel
	PushToChannel(chanelId int64, message interface{}) (err error)

	IsConnected(userId int64)
	ConnectionHandle(handler ConnectionHandler)
}

type ConnectionHandler interface {
	Connected(userId int64)
	Disconnected(userId int64)
}
