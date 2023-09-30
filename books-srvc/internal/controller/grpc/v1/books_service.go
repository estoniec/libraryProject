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
