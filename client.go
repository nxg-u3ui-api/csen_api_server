package main

import (
	"log"
	"github.com/gorilla/websocket"
)

type client struct {
	socket *websocket.Conn
	send chan []byte
}

func (c *client) read() {
	for {
		if _, msg, err := c.socket.ReadMessage(); err == nil {
			// something to do
			log.Print(msg)
		} else {
			break
		}
	}
	c.socket.Close()
}

func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
