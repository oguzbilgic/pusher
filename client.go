// Package pusher implements client library for pusher.com socket
package pusher

import (
	"code.google.com/p/go.net/websocket"
	"fmt"
	"time"
)

const (
	pusherUrl = "ws://ws.pusherapp.com:80/app/%s?protocol=7"
)

type Client struct {
	key      string
	conn     *websocket.Conn
	channels []*Channel
}

func New(key string) (*Client, error) {
	ws, err := websocket.Dial(fmt.Sprintf(pusherUrl, key), "", "http://localhost/")
	if err != nil {
		return nil, err
	}

	client := &Client{
		key:  key,
		conn: ws,
		channels: []*Channel{
			NewChannel(""),
		},
	}

	go client.pong()
	go client.poll()

	return client, nil
}

func (c *Client) pong() {
	tick := time.Tick(time.Minute)
	pong := NewPongMessage()
	for {
		<-tick
		websocket.JSON.Send(c.conn, pong)
	}
}

func (c *Client) poll() {
	for {
		var msg Message
		err := websocket.JSON.Receive(c.conn, &msg)
		if err != nil {
			panic(err)
		}

		c.processMessage(&msg)
	}
}

func (c *Client) processMessage(msg *Message) {
	for _, channel := range c.channels {
		if channel.Name == msg.Channel {
			channel.processMessage(msg)
		}
	}
}

func (c *Client) Disconnect() error {
	return c.conn.Close()
}

func (c *Client) Channel(name string) *Channel {
	for _, channel := range c.channels {
		if channel.Name == name {
			return channel
		}
	}

	channel := NewChannel(name)
	c.channels = append(c.channels, channel)
	websocket.JSON.Send(c.conn, NewSubscribeMessage(name))

	return channel
}
