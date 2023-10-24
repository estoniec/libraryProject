package v1

import (
	"context"
	pb "github.com/estoniec/libraryProject/contracts/gen/go/books_users"
	"rent/internal/domain/rent/dto"
)

type Service interface {
	RentBook(ctx context.Context, input dto.RentBookInput) (dto.RentBookOutput, error)
}

type Server struct {
	service Service
	pb.UnimplementedBooksUsersServiceServer
}

func NewServer(
	service Service,
	srv pb.UnimplementedBooksUsersServiceServer,
) *Server {
	return &Server{
		service:                              service,
		UnimplementedBooksUsersServiceServer: srv,
	}
}