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

func (s *Server) GetDebt(ctx context.Context, req *pb.GetDebtRequest) (*pb.GetDebtResponse, error) {
	dto := NewGetDebtInput(req)

	res, err := s.service.GetDebt(ctx, dto)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	response := NewGetDebtOutput(res)

	return response, nil
}

func (s *Server) FindByUidAndBid(ctx context.Context, req *pb.FindByUidAndBidRequest) (*pb.FindByUidAndBidResponse, error) {
	dto := NewFindByUidAndBidInput(req)

	res, err := s.service.FindByUidAndBid(ctx, dto)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	response := NewFindByUidAndBidOutput(res)

	return response, nil
}

func (s *Server) ConfirmReturn(ctx context.Context, req *pb.ConfirmReturnRequest) (*pb.ConfirmReturnResponse, error) {
	dto := NewConfirmReturnInput(req)

	res, err := s.service.ConfirmReturn(ctx, dto)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	response := NewConfirmReturnOutput(res)

	return response, nil
}

func (s *Server) FindBy(ctx context.Context, req *pb.FindByRequest) (*pb.FindByResponse, error) {
	dto := NewFindByInput(req)

	res, err := s.service.FindBy(ctx, dto)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	response := NewFindByOutput(res)

	return response, nil
}
