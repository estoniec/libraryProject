package v1

import (
	"context"
	pb "github.com/estoniec/automaticLibrary/contracts/gen/go/registration"
	"log/slog"
	dto2 "project11/registration-svc/internal/domain/reg/dto"
)

func (s *Server) Registration(ctx context.Context, req *pb.RegRequest) (*pb.RegResponse, error) {
	user := NewRegInput(req)

	res, err := s.service.RegUser(ctx, user)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	response := NewRegOutput(res)

	return response, nil
}

func (s *Server) CheckUser(ctx context.Context, req *pb.CheckRequest) (*pb.CheckResponse, error) {
	user := NewCheckInput(req)

	dto := dto2.NewCheckInput(user.ID)

	res, err := s.service.CheckUser(ctx, dto)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	response := NewCheckOutput(res)

	return response, nil
}
