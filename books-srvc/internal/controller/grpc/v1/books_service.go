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
