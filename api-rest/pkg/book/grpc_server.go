package book

import (
	"context"
	"fmt"
	"log"
	"projeto-go/api-rest/pkg/models"
	"strconv"

	"gorm.io/gorm"

	pb "projeto-go/api-rest/internal/grpc"
)

type Server struct {
	pb.UnimplementedApiBookServiceServer
	DB *gorm.DB
}

func (s *Server) AddBook(ctx context.Context, req *pb.Book) (*pb.AddBookResponse, error) {
	book := models.Book{
		Author: req.GetAuthor(),
		Name:   req.GetName(),
		Price:  req.GetPrice(),
	}

	createdBook, err := AddBook(s.DB, book)
	if err != nil {
		return nil, err
	}

	grpcAddBook := &pb.AddBookResponse{
		Id:     strconv.Itoa(int(createdBook.ID)),
		Author: createdBook.Author,
		Name:   createdBook.Name,
		Price:  createdBook.Price,
	}

	message := fmt.Sprintf("Novo livro adicionado: %s por %s", createdBook.Name, createdBook.Author)
	log.Printf("Sending message: %s", message)
	if grpcWebSocketClient != nil {
		_, err = grpcWebSocketClient.BroadcastMessage(ctx, &pb.BroadcastRequest{
			Event: "new_book_added",
			Data:  message,
		})
		if err != nil {
			log.Printf("Send message via gRPC error: %v", err)
		}
	} else {
		log.Println("gRPC WebSocket client is not initialized.")
	}

	return grpcAddBook, nil
}
