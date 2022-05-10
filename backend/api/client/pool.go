package client

import (
	"fmt"
)

type Pool struct {
	Register          chan *Client
	Unregister        chan *Client
	Clients           map[*Client]bool
	PuzzleUpdate      chan *Client
	SolvePuzzleStart  chan *Client
	SolvePuzzleStop   chan bool
	SolvePuzzleResume chan bool
	PuzzleSolved      chan *Client
}

func NewPool() *Pool {
	return &Pool{
		Register:          make(chan *Client),
		Unregister:        make(chan *Client),
		Clients:           make(map[*Client]bool),
		PuzzleUpdate:      make(chan *Client),
		SolvePuzzleStart:  make(chan *Client),
		SolvePuzzleStop:   make(chan bool),
		SolvePuzzleResume: make(chan bool),
		PuzzleSolved:      make(chan *Client),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			fmt.Printf("Size of Connection Pool: %d\n", len(pool.Clients))

			arr := GenerateRandomArray(10)
			client.Puzzle = arr
			messageBody := ConvertSliceToString(arr)

			if err := client.Conn.WriteMessage(1, []byte(messageBody)); err != nil {
				fmt.Printf("Error in Puzzle update: %v\n", err)
				return
			}

			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			fmt.Printf("Size of Connection Pool: %d\n", len(pool.Clients))

			if err := client.Conn.WriteMessage(1, []byte("0:0")); err != nil {
				fmt.Printf("Error in Puzzle update: %v\n", err)
				return
			}
			break
		case client := <-pool.PuzzleUpdate:
			arr := client.Puzzle
			messageBody := ConvertSliceToString(arr)

			if err := client.Conn.WriteMessage(1, []byte(messageBody)); err != nil {
				fmt.Printf("Error in Puzzle update: %v\n", err)
				return
			}
			break
		case client := <-pool.SolvePuzzleStart:
			arr := client.Puzzle

			go BinarySort(client, 1, arr)
			break
		case client := <-pool.PuzzleSolved:
			arr := client.Puzzle
			messageBody := ConvertSliceToString(arr)

			if err := client.Conn.WriteMessage(1, []byte("solved:"+messageBody)); err != nil {
				fmt.Printf("Error in Puzzle update: %v\n", err)
				return
			}
			delete(pool.Clients, client)
			fmt.Printf("Size of Connection Pool: %d\n", len(pool.Clients))
		}

	}
}
