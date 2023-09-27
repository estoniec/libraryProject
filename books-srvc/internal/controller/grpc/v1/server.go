package v1

import (
	"books-srvc/internal/domain/books/dto"
	"context"
	pb "github.com/estoniec/libraryProject/contracts/gen/go/books"
)

type Service interface {
	FindByISBN(ctx context.Context, dto dto.FindByISBNInput) (dto.FindByISBNOutput, error)
}

type Server struct {
	service Service
	pb.UnimplementedBooksServiceServer
}

func NewServer(
	service Service,
	srv pb.UnimplementedBooksServiceServer,
) *Server {
	return &Server{
		service:                         service,
		UnimplementedBooksServiceServer: srv,
	}
}
