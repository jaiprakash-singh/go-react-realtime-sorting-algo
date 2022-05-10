package routes

import (
	"fmt"
	"net/http"

	"github.com/jaiprakash-singh/go-react-realtime-sorting-algo/api/client"
	"github.com/jaiprakash-singh/go-react-realtime-sorting-algo/api/server"
)

func SetupRoutes() {
	pool := client.NewPool()
	go pool.Start()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to Web Socket API Home!")
	})

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		server.ServeWS(pool, w, r)
	})
}
