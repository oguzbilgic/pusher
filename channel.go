package pusher

type Channel struct {
	Name  string
	binds map[string][]chan *Message
}

func NewChannel(name string) *Channel {
	return &Channel{name, make(map[string][]chan *Message)}
}

func (c *Channel) Bind(event string) chan *Message {
	eventChan := make(chan *Message)
	c.binds[event] = append(c.binds[event], eventChan)
	return eventChan
}

func (c *Channel) processMessage(msg *Message) {
	for _, eventChan := range c.binds[msg.Event] {
		eventChan <- msg
	}
}
