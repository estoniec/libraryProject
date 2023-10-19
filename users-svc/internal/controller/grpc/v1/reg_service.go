package v1

import (
	"context"
	pb "github.com/estoniec/libraryProject/contracts/gen/go/users"
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

func (s *Server) CheckRole(ctx context.Context, req *pb.CheckRoleRequest) (*pb.CheckRoleResponse, error) {
	dto := NewCheckRoleInput(req)

	res, err := s.service.CheckRole(ctx, dto)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}

	response := NewCheckRoleOutput(res)

	return response, nil
}
