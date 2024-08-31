package book

import (
	"log"
	pb "projeto-go/api-rest/internal/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var grpcWebSocketClient pb.WebSocketServiceClient

func InitWebSocketGRPCClient(address string) error {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to websocket gRPC server: %v", err)
	}
	grpcWebSocketClient = pb.NewWebSocketServiceClient(conn)
	return nil
}
