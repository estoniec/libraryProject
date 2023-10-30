package v1

import (
	"context"
	pb "github.com/estoniec/libraryProject/contracts/gen/go/books_users"
	"rent/internal/domain/rent/dto"
)

type Service interface {
	RentBook(ctx context.Context, input dto.RentBookInput) (dto.RentBookOutput, error)
	FindBook(ctx context.Context, input dto.FindBookInput) (dto.FindBookOutput, error)
	ConfirmRent(ctx context.Context, input dto.ConfirmRentInput) (dto.ConfirmRentOutput, error)
	GetDebt(ctx context.Context, input dto.GetDebtInput) (dto.GetDebtOutput, error)
	FindByUidAndBid(ctx context.Context, input dto.FindByUIDAndBIDInput) (dto.FindByUIDAndBIDOutput, error)
	ConfirmReturn(ctx context.Context, input dto.ConfirmReturnInput) (dto.ConfirmReturnOutput, error)
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
