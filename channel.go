package pusher

type Channel struct {
	Name      string
	dataChans map[string][]chan interface{}
}

func NewChannel(name string) *Channel {
	return &Channel{name, make(map[string][]chan interface{})}
}

func (c *Channel) Bind(event string) chan interface{} {
	dataChan := make(chan interface{})
	c.dataChans[event] = append(c.dataChans[event], dataChan)
	return dataChan
}

func (c *Channel) processMessage(msg *Message) {
	for _, dataChan := range c.dataChans[msg.Event] {
		dataChan <- msg.Data
	}
}
