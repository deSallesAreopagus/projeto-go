package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"websocket-go/internal/grpc"
	"websocket-go/pkg/handlers"
	"websocket-go/pkg/services"

	g "google.golang.org/grpc"
)

func main() {
	webSocketServer := services.NewWebSocketServer()

	go func() {
		http.HandleFunc("/ws", handlers.WSHandler(webSocketServer))

		log.Println("WebSocket server started on :8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

	grpcServer := g.NewServer()
	grpc.RegisterWebSocketServiceServer(grpcServer, webSocketServer)

	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen on port 50052: %v", err)
	}

	fmt.Println("Starting gRPC WebSocket server on port 50052...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}
