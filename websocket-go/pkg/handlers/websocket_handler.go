package handlers

import (
	"log"
	"net/http"

	"websocket-go/pkg/services"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]struct{})

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WSHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, "Websocket connection error", http.StatusBadRequest)
			return
		}

		clients[conn] = struct{}{}

		go listenToClient(conn)
	}
}

func listenToClient(conn *websocket.Conn) {
	defer func() {
		delete(clients, conn)
		conn.Close()
	}()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error: ", err)
			break
		}

		response := services.HandleMessage(string(message))

		if err := conn.WriteMessage(messageType, []byte(response)); err != nil {
			log.Println("Write error: ", err)
			break
		}
	}
}
