package server

import (
	"fmt"
	"net/http"

	"github.com/jaiprakash-singh/go-react-realtime-sorting-algo/api/client"

	"github.com/google/uuid"
)

//define our WebSocket endpoint
func ServeWS(pool *client.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Printf("WebSocket Endpoint Reached!\n")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}

	client := &client.Client{
		ID:   uuid.New().String(),
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}
