# Pusher [![Build Status](https://travis-ci.org/oguzbilgic/pusher.png?branch=master)](https://travis-ci.org/oguzbilgic/pusher)

Package pusher implements client library for [pusher.com](http://pusher.com/docs/)

## Usage

```go
package main

import (
	"github.com/oguzbilgic/pusher"
)

func main() {
	conn, err := pusher.New("d05049c57n3ielfhfh82")
	if err != nil {
		panic(err)
	}

	chatRoomChan := conn.Channel("chat_room")
	messages := chatRoomChan.Bind("new_message")

	for {
		msg := <-messages

		println(msg)
	}
}
```

## Documentation 

http://godoc.org/github.com/oguzbilgic/pusher

## License

The MIT License (MIT)
