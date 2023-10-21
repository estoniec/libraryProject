package v1

import (
	"context"
	pb "github.com/estoniec/libraryProject/contracts/gen/go/books_users"
	"log/slog"
)

func (s *Server) RentBook(ctx context.Context, req *pb.RentBookRequest) (*pb.RentBookResponse, error) {
	dto := NewRentBookInput(req)

	res, err := s.service.RentBook(ctx, dto)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	response := NewRentBookOutput(res)

	return response, nil
}
