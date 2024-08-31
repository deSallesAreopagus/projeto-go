package services

import (
	"context"
	"fmt"
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

func (s *WebSocketServer) BroadcastMessage(ctx context.Context, req *grpc.BroadcastRequest) (*grpc.BroadcastResponse, error) {
	if len(s.clients) == 0 {
		log.Printf("No clients connected. Clients: %v", s.clients)
		return &grpc.BroadcastResponse{Status: "fail"}, nil
	}

	message := req.GetData()

	for client := range s.clients {
		err := client.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			log.Printf("Failed to send message to client: %v", err)

		}
	}

	return &grpc.BroadcastResponse{Status: "success"}, nil
}

func HandleMessage(message string) string {
	response := "Message sent: " + message

	fmt.Printf("Message sent: %s\n", message)

	return response
}
