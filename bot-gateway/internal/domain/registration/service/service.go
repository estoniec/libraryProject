package service

import (
	"context"
	"gateway/internal/domain/registration/dto"
	pb "github.com/estoniec/automaticLibrary/contracts/gen/go/registration"
)

type Service struct {
	client pb.RegServiceClient
}

func NewService(client pb.RegServiceClient) *Service {
	return &Service{
		client: client,
	}
}

func (s *Service) Registration(ctx context.Context, input dto.RegInput) (dto.RegOutput, error) {
	res, err := s.client.Registration(ctx, &pb.RegRequest{
		ID:       input.ID,
		Username: input.Username,
		Class:    input.Class,
		Phone:    input.Phone,
	})
	if err != nil {
		response := dto.NewRegOutput(err.Error(), 404)
		return response, err
	}
	response := dto.NewRegOutput(res.GetError(), res.GetStatus())
	return response, nil
}

func (s *Service) CheckUser(ctx context.Context, input dto.CheckInput) (dto.CheckOutput, error) {
	res, err := s.client.CheckUser(ctx, &pb.CheckRequest{
		ID: input.ID,
	})
	if err != nil {
		response := dto.NewCheckOutput(false)
		return response, err
	}
	response := dto.NewCheckOutput(res.GetChecked())
	return response, nil
}
