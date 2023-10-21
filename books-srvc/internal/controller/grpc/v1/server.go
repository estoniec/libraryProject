package v1

import (
	"books-srvc/internal/domain/books/dto"
	"context"
	pb "github.com/estoniec/libraryProject/contracts/gen/go/books"
)

type Service interface {
	FindBy(ctx context.Context, input dto.FindByInput) (dto.FindByOutput, error)
	CreateBook(ctx context.Context, input dto.CreateBookInput) (dto.CreateBookOutput, error)
	EditCountBook(ctx context.Context, input dto.EditCountBookInput) (dto.EditCountBookOutput, error)
	DeleteBook(ctx context.Context, input dto.DeleteBookInput) (dto.DeleteBookOutput, error)
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
