package v1

import (
	"context"
	pb "github.com/estoniec/automaticLibrary/contracts/gen/go/registration"
	"project11/registration-svc/internal/domain/reg/dto"
	"project11/registration-svc/internal/domain/reg/model"
)

type Service interface {
	RegUser(ctx context.Context, model model.User) (dto.RegOutput, error)
	CheckUser(ctx context.Context, input dto.CheckInput) (dto.CheckOutput, error)
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
