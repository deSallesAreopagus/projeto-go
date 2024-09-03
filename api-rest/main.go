package main

import (
	"fmt"
	"log"
	"net"
	"os"
	pb "projeto-go/api-rest/internal/grpc"
	"projeto-go/api-rest/pkg/book"
	"projeto-go/api-rest/pkg/database"
	"projeto-go/api-rest/pkg/kafka"
	"projeto-go/api-rest/pkg/router"

	"google.golang.org/grpc"
)

func main() {
	websocketString := os.Getenv("WEBSOCKET_STRING")
	if websocketString == "" {
		websocketString = "localhost:50052"
	}
	err := book.InitWebSocketGRPCClient(websocketString)
	if err != nil {
		log.Fatalf("Failed to initialize WebSocket gRPC client: %v", err)
	}
	fmt.Println("Connected to gRPC client on port :50052...")

	kafkaBroker := os.Getenv("KAFKA_BROKER")
	if kafkaBroker == "" {
		kafkaBroker = "localhost:9092"
	}
	kafka.InitProducer(kafkaBroker)
	defer kafka.CloseProducer()

	db := database.SetupDBConnection()
	defer database.CloseDBConnection(db)

	r := router.SetupRouter(db)
	go func() {
		if err := r.Run(":8000"); err != nil {
			log.Fatalf("Failed to start HTTP server: %v", err)
		}
	}()

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}
	grpcServer := grpc.NewServer()
	bookServer := &book.Server{DB: db}
	pb.RegisterApiBookServiceServer(grpcServer, bookServer)

	fmt.Println("Starting gRPC server on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}
}
