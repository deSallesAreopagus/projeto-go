package main

import (
	"log"
	"net/http"

	"websocket-go/internal/handlers"
)

func main() {
	http.HandleFunc("/ws", handlers.WSHandler())

	log.Println("WebSocket server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
