package client

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID     string
	Conn   *websocket.Conn
	Pool   *Pool
	Puzzle []int
}

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	for {
		messageType, p, err := c.Conn.ReadMessage()
		if err != nil {
			log.Printf("Error: %v\n", err)
		}

		if string(p) == "sort" {
			c.Pool.SolvePuzzleStart <- c
		} else if string(p) == "stop" {
			c.Pool.SolvePuzzleStop <- true
		} else if string(p) == "resume" {
			c.Pool.SolvePuzzleResume <- true
		}

		fmt.Printf("Message: %d:%s\n", messageType, string(p))
	}
}
