// Package pusher implements client library for pusher.com socket
package pusher

type Client struct {
	appKey string
}

func New(appKey string) *Client {
	return &Client{appKey}
}

func (c *Client) Channel(name string) *Channel {
	return &Channel{name}
}
