package main

import (
	"fmt"
	"net/http"

	"github.com/nek07/go-chat/backend/pkg/websocket"
)

func serveWS(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {

}
func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWS(pool, w, r)
	})
}
func main() {
	fmt.Println("Hello world!")
	setupRoutes()
	http.ListenAndServe(":9000", nil)
}
