// Package pusher implements client library for pusher.com socket
package pusher

import (
	"fmt"
	"time"

	"golang.org/x/net/websocket"
)

const (
	pusherUrl = "ws://ws.pusherapp.com:80/app/%s?protocol=7"
)

type Connection struct {
	key      string
	conn     *websocket.Conn
	channels []*Channel
}

func New(key string) (*Connection, error) {
	return NewFromKeyAndUrl(key, pusherUrl)
}

func NewFromKeyAndUrl(key string, keylessUrl string) (*Connection, error) {
	ws, err := websocket.Dial(
		fmt.Sprintf(keylessUrl, key), "", "http://localhost/")
	if err != nil {
		return nil, err
	}

	connection := &Connection{
		key:  key,
		conn: ws,
		channels: []*Channel{
			NewChannel(""),
		},
	}

	go connection.pong()
	go connection.poll()

	return connection, nil
}

func (c *Connection) pong() {
	tick := time.Tick(time.Minute)
	pong := NewPongMessage()
	for {
		<-tick
		websocket.JSON.Send(c.conn, pong)
	}
}

func (c *Connection) poll() {
	for {
		var msg Message
		err := websocket.JSON.Receive(c.conn, &msg)
		if err != nil {
			panic(err)
		}

		c.processMessage(&msg)
	}
}

func (c *Connection) processMessage(msg *Message) {
	for _, channel := range c.channels {
		if channel.Name == msg.Channel {
			channel.processMessage(msg)
		}
	}
}

func (c *Connection) Disconnect() error {
	return c.conn.Close()
}

func (c *Connection) Channel(name string) *Channel {
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
