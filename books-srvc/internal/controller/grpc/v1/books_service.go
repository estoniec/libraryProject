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
