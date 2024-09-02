package services

import (
	"context"
	"log"
	"websocket-go/internal/grpc"

	"github.com/gorilla/websocket"
)

type WebSocketServer struct {
	grpc.UnimplementedWebSocketServiceServer
	clients map[*websocket.Conn]struct{}
}

func NewWebSocketServer() *WebSocketServer {
	return &WebSocketServer{
		clients: make(map[*websocket.Conn]struct{}),
	}
}

func (s *WebSocketServer) RegisterClient(conn *websocket.Conn) {
	s.clients[conn] = struct{}{}
}

func (s *WebSocketServer) UnregisterClient(conn *websocket.Conn) {
	delete(s.clients, conn)
}

func (s *WebSocketServer) BroadcastMessage(ctx context.Context, req *grpc.BroadcastRequest) (*grpc.BroadcastResponse, error) {
	if len(s.clients) == 0 {
		log.Printf("No clients connected. Clients: %v", s.clients)
		return &grpc.BroadcastResponse{Status: "fail"}, nil
	}

	message := req.GetData()

	log.Print("Clientes conectados: ", len(s.clients))
	for client := range s.clients {
		err := client.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Printf("Failed to send message to client: %v", err)
			s.UnregisterClient(client)
		}
	}

	return &grpc.BroadcastResponse{Status: "success"}, nil
}

func HandleMessage(message string) string {
	response := "Message sent: " + message

	log.Printf("Message sent: %s\n", message)

	return response
}
