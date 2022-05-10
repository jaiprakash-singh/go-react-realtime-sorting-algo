package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jaiprakash-singh/go-react-realtime-sorting-algo/routes"
)

func main() {
	fmt.Printf("Realtime ticker App\n")

	routes.SetupRoutes()

	log.Fatal(http.ListenAndServe(":9001", nil))

	//arr := []int{2, 5, 6, 7, 9, 1, 3}
	//fmt.Printf("Result: %v\n", arr)
	//fmt.Printf("Result: %v\n", client.BinarySort(arr))
}
