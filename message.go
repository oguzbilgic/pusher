package pusher

type Message struct {
	Event   string      `json:"event"`
	Channel string      `json:"channel"`
	Data    interface{} `json:"data"`
}

func NewSubscribeMessage(channel string) *Message {
	return &Message{"pusher:subscribe", "", map[string]string{"channel": channel}}
}
