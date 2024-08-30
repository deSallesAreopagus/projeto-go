package services

import (
	"context"
	"encoding/json"
	"log"
	"strings"

	pb "websocket-go/internal/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func HandleMessage(message string) string {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewBookServiceClient(conn)

	switch strings.ToLower(message) {
	case "ler dados":
		return fetchBooks(client)
	default:
		return addBook(client, message)
	}
}

func fetchBooks(client pb.BookServiceClient) string {
	res, err := client.GetBooks(context.Background(), &pb.GetBooksRequest{})
	if err != nil {
		return "Error fetching books"
	}

	booksJson, err := json.Marshal(res.Books)
	if err != nil {
		return "Error converting to JSON"
	}

	return string(booksJson)
}

func addBook(client pb.BookServiceClient, author string) string {
	book := &pb.Book{Author: author, Name: "Teste", Price: "59.99"}
	_, err := client.AddBook(context.Background(), book)
	if err != nil {
		return "Error saving to database"
	}
	return "Book added: " + author
}
