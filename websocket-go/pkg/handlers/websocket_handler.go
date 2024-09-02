package handlers

import (
	"log"
	"net/http"
	"websocket-go/pkg/services"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WSHandler(server *services.WebSocketServer) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			http.Error(w, "Websocket connection error", http.StatusBadRequest)
			return
		}

		server.RegisterClient(conn)

		go listenToClient(server, conn)
	}
}

func listenToClient(server *services.WebSocketServer, conn *websocket.Conn) {
	defer func() {
		server.UnregisterClient(conn)
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
