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
func (s *Server) UpdateBook(ctx context.Context, req *pb.UpdateBookRequest) (*pb.UpdateBookResponse, error) {
	bookID, err := strconv.Atoi(req.GetId())
	if err != nil {
		return nil, fmt.Errorf("invalid ID: %v", err)
	}

	updatedBook, err := ModifyBook(s.DB, uint(bookID), models.Book{
		Author: req.GetAuthor(),
		Name:   req.GetName(),
		Price:  req.GetPrice(),
	})
	if err != nil {
		return nil, err
	}

	grpcUpdateBook := &pb.UpdateBookResponse{
		Id:     strconv.Itoa(int(updatedBook.ID)),
		Author: updatedBook.Author,
		Name:   updatedBook.Name,
		Price:  updatedBook.Price,
	}

	message := fmt.Sprintf("Livro atualizado: %s por %s", updatedBook.Name, updatedBook.Author)
	log.Printf("Sending message: %s", message)
	if grpcWebSocketClient != nil {
		_, err = grpcWebSocketClient.BroadcastMessage(ctx, &pb.BroadcastRequest{
			Event: "book_updated",
			Data:  message,
		})
		if err != nil {
			log.Printf("Send message via gRPC error: %v", err)
		}
	} else {
		log.Println("gRPC WebSocket client is not initialized.")
	}

	return grpcUpdateBook, nil
}
func (s *Server) DeleteBook(ctx context.Context, req *pb.DeleteBookRequest) (*pb.DeleteBookResponse, error) {
	bookID, err := strconv.Atoi(req.GetId())
	if err != nil {
		return nil, fmt.Errorf("invalid ID: %v", err)
	}

	err = RemoveBook(s.DB, uint(bookID))
	if err != nil {
		return nil, err
	}

	grpcDeleteBook := &pb.DeleteBookResponse{
		Id: req.GetId(),
	}

	message := fmt.Sprintf("Livro deletado: ID - %s", req.GetId())
	log.Printf("Sending message: %s", message)
	if grpcWebSocketClient != nil {
		_, err = grpcWebSocketClient.BroadcastMessage(ctx, &pb.BroadcastRequest{
			Event: "book_deleted",
			Data:  message,
		})
		if err != nil {
			log.Printf("Send message via gRPC error: %v", err)
		}
	} else {
		log.Println("gRPC WebSocket client is not initialized.")
	}

	return grpcDeleteBook, nil
}
func (s *Server) GetBooks(ctx context.Context, req *pb.GetBooksRequest) (*pb.GetBooksResponse, error) {
	books, err := FindAllBooks(s.DB)
	if err != nil {
		return nil, err
	}

	var grpcBooks []*pb.Book
	for _, book := range books {
		grpcBooks = append(grpcBooks, &pb.Book{
			Id:     strconv.Itoa(int(book.ID)),
			Author: book.Author,
			Name:   book.Name,
			Price:  book.Price,
		})
	}

	return &pb.GetBooksResponse{Books: grpcBooks}, nil
}
func (s *Server) GetBookById(ctx context.Context, req *pb.GetBookByIdRequest) (*pb.GetBookByIdResponse, error) {
	bookID, err := strconv.Atoi(req.GetId())
	if err != nil {
		return nil, fmt.Errorf("invalid ID: %v", err)
	}

	book, err := FindBookById(s.DB, uint(bookID))
	if err != nil {
		return nil, err
	}

	return &pb.GetBookByIdResponse{
		Book: &pb.Book{
			Id:     strconv.Itoa(int(book.ID)),
			Author: book.Author,
			Name:   book.Name,
			Price:  book.Price,
		},
	}, nil
}
