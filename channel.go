package pusher

type Channel struct {
	name string
}

func (c *Channel) Bind(eventName string) chan *Event {
	return nil
}
