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

func (s *Server) FindBook(ctx context.Context, req *pb.FindBookRequest) (*pb.FindBookResponse, error) {
	dto := NewFindBookInput(req)

	res, err := s.service.FindBook(ctx, dto)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	response := NewFindBookOutput(res)

	return response, nil
}

func (s *Server) ConfirmRent(ctx context.Context, req *pb.ConfirmRentRequest) (*pb.ConfirmRentResponse, error) {
	dto := NewConfirmRentInput(req)

	res, err := s.service.ConfirmRent(ctx, dto)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	response := NewConfirmRentOutput(res)

	return response, nil
}
