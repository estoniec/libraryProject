package v1

import (
	"books-srvc/internal/domain/books/dto"
	"context"
	pb "github.com/estoniec/libraryProject/contracts/gen/go/books"
)

type Service interface {
	FindByISBN(ctx context.Context, dto dto.FindByISBNInput) (dto.FindByISBNOutput, error)
	FindByAuthor(ctx context.Context, dto dto.FindByAuthorInput) (dto.FindByAuthorOutput, error)
	FindByName(ctx context.Context, dto dto.FindByNameInput) (dto.FindByNameOutput, error)
	FindByNameAndAuthor(ctx context.Context, dto dto.FindByNameAndAuthorInput) (dto.FindByNameAndAuthorOutput, error)
	FindAll(ctx context.Context, dto dto.FindAllInput) (dto.FindAllOutput, error)
	FindBy(ctx context.Context, input dto.FindByInput) (dto.FindByOutput, error)
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
