package v1

import (
	"context"
	pb "github.com/estoniec/libraryProject/contracts/gen/go/registration"
	"log/slog"
)

func (s *Server) Registration(ctx context.Context, req *pb.RegRequest) (*pb.RegResponse, error) {
	dto := NewRegInput(req)

	res, err := s.service.RegUser(ctx, dto)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	response := NewRegOutput(res)

	return response, nil
}

func (s *Server) CheckUser(ctx context.Context, req *pb.CheckRequest) (*pb.CheckResponse, error) {
	dto := NewCheckInput(req)

	res, err := s.service.CheckUser(ctx, dto)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	response := NewCheckOutput(res)

	return response, nil
}
