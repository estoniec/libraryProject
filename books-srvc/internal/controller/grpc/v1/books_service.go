package v1

import (
	"context"
	pb "github.com/estoniec/libraryProject/contracts/gen/go/books"
	"log/slog"
)

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

func (s *Server) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.CreateBookResponse, error) {
	dto := NewCreateInput(req)

	res, err := s.service.CreateBook(ctx, dto)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	response := NewCreateOutput(res)

	return response, nil
}

func (s *Server) EditCountBook(ctx context.Context, req *pb.EditCountBookRequest) (*pb.EditCountBookResponse, error) {
	dto := NewEditCountInput(req)

	res, err := s.service.EditCountBook(ctx, dto)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	response := NewEditCountOutput(res)

	return response, nil
}

func (s *Server) DeleteBook(ctx context.Context, req *pb.DeleteBookRequest) (*pb.DeleteBookResponse, error) {
	dto := NewDeleteInput(req)

	res, err := s.service.DeleteBook(ctx, dto)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	response := NewDeleteOutput(res)

	return response, nil
}
