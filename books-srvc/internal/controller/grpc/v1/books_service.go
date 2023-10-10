package v1

import (
	"context"
	pb "github.com/estoniec/libraryProject/contracts/gen/go/books"
	"log/slog"
)

func (s *Server) FindByISBN(ctx context.Context, req *pb.FindByISBNRequest) (*pb.FindByISBNResponse, error) {
	dto := NewFindByISBNInput(req)

	res, err := s.service.FindByISBN(ctx, dto)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	response := NewFindByISBNOutput(res)

	return response, nil
}

func (s *Server) FindByAuthor(ctx context.Context, req *pb.FindByAuthorRequest) (*pb.FindByAuthorResponse, error) {
	dto := NewFindByAuthorInput(req)

	res, err := s.service.FindByAuthor(ctx, dto)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	response := NewFindByAuthorOutput(res)

	return response, nil
}

func (s *Server) FindByName(ctx context.Context, req *pb.FindByNameRequest) (*pb.FindByNameResponse, error) {
	dto := NewFindByNameInput(req)

	res, err := s.service.FindByName(ctx, dto)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	response := NewFindByNameOutput(res)

	return response, nil
}

func (s *Server) FindByNameAndAuthor(ctx context.Context, req *pb.FindByNameAndAuthorRequest) (*pb.FindByNameAndAuthorResponse, error) {
	dto := NewFindByNameAndAuthorInput(req)

	res, err := s.service.FindByNameAndAuthor(ctx, dto)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	response := NewFindByNameAndAuthorOutput(res)

	return response, nil
}

func (s *Server) FindAll(ctx context.Context, req *pb.FindAllRequest) (*pb.FindAllResponse, error) {
	dto := NewFindAllInput(req)

	res, err := s.service.FindAll(ctx, dto)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	response := NewFindAllOutput(res)

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
