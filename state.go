package pusher

type State int

const (
	Initialized State = iota
	Connecting
	Connected
	Unavailable
	Failed
	Disconnected
)
