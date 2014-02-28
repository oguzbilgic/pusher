package pusher

type Channel struct {
	name string
}

func (c *Channel) Bind(event string) chan *Message {
	return nil
}
