package book

import (
	"context"
	"log"
	"projeto-go/api-rest/pkg/models"
	"strconv"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	pb "projeto-go/api-rest/internal/grpc"
)

type Server struct {
	pb.UnimplementedApiBookServiceServer
	DB *gorm.DB
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
		return nil, status.Errorf(codes.InvalidArgument, "Invalid ID")
	}

	book, err := FindBookById(s.DB, uint(bookID))

	if err != nil {
		return nil, status.Error(codes.Internal, "Internal server error")
	}

	return &pb.GetBookByIdResponse{Book: &pb.Book{
		Id:     strconv.Itoa(int(book.ID)),
		Author: book.Author,
		Name:   book.Name,
		Price:  book.Price,
	}}, nil
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

	grpcAddBook := &pb.AddBookResponse{}

	grpcAddBook.Id = strconv.Itoa(int(createdBook.ID))
	grpcAddBook.Author = createdBook.Author
	grpcAddBook.Name = createdBook.Name
	grpcAddBook.Price = createdBook.Price

	log.Printf("Book: %+v", grpcAddBook)

	return grpcAddBook, nil
}

func (s *Server) UpdateBook(ctx context.Context, req *pb.UpdateBookRequest) (*pb.UpdateBookResponse, error) {
	var book models.Book

	bookID, err := strconv.Atoi(req.GetId())

	if err != nil {
		return nil, err
	}

	book.Author = req.GetAuthor()
	book.Name = req.GetName()
	book.Price = req.GetPrice()

	book, err = ModifyBook(s.DB, uint(bookID), book)

	if err != nil {
		return nil, err
	}

	return &pb.UpdateBookResponse{
		Id:     strconv.Itoa(int(book.ID)),
		Author: book.Author,
		Name:   book.Name,
		Price:  book.Price,
	}, nil
}

func (s *Server) DeleteBook(ctx context.Context, req *pb.DeleteBookRequest) (*pb.DeleteBookResponse, error) {
	bookID, err := strconv.Atoi(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid ID")
	}

	err = RemoveBook(s.DB, uint(bookID))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Internal server error")
	}

	return &pb.DeleteBookResponse{
		Id: req.GetId(),
	}, nil
}
