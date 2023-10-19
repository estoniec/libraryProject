package v1

import (
	"context"
	pb "github.com/estoniec/libraryProject/contracts/gen/go/users"
	"registration-svc/internal/domain/reg/dto"
	"registration-svc/internal/domain/reg/model"
)

type Service interface {
	RegUser(ctx context.Context, model model.User) (dto.RegOutput, error)
	CheckUser(ctx context.Context, input dto.CheckInput) (dto.CheckOutput, error)
	CheckRole(ctx context.Context, input dto.CheckRoleInput) (dto.CheckRoleOutput, error)
}

type Server struct {
	service Service
	pb.UnimplementedRegServiceServer
}

func NewServer(
	service Service,
	srv pb.UnimplementedRegServiceServer,
) *Server {
	return &Server{
		service:                       service,
		UnimplementedRegServiceServer: srv,
	}
}
