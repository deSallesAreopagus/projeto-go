package main

import (
	"fmt"
	"log"
	"net"
	pb "projeto-go/api-rest/internal/grpc"
	"projeto-go/api-rest/pkg/book"
	"projeto-go/api-rest/pkg/database"
	"projeto-go/api-rest/pkg/kafka"
	"projeto-go/api-rest/pkg/router"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	kafka.InitProducer("localhost:9092")
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
