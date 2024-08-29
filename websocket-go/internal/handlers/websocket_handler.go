package handlers

import (
	"log"
	"net/http"

	"websocket-go/internal/services"
	"websocket-go/pkg/websocket"
)

func WSHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := websocket.Upgrade(w, r)
		if err != nil {
			log.Println("Upgrade error:", err)
			return
		}
		defer conn.Close()

		for {
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				log.Println("Read error:", err)
				break
			}

			response := services.HandleMessage(string(message))

			if err := conn.WriteMessage(messageType, []byte(response)); err != nil {
				log.Println("Write error:", err)
				break
			}
		}
	}
}
