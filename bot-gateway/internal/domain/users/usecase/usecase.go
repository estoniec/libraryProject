package users_usecase

import (
	"context"
	dto2 "gateway/internal/controller/telegram/dto"
	dto "gateway/internal/domain/users/dto"
	pb "github.com/estoniec/libraryProject/contracts/gen/go/users"
)

type Usecase struct {
	client pb.RegServiceClient
}

func NewUsecase(client pb.RegServiceClient) *Usecase {
	return &Usecase{
		client: client,
	}
}

func (s *Usecase) Registration(ctx context.Context, input dto2.RegInput) (dto.RegOutput, error) {
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

func (s *Usecase) CheckUser(ctx context.Context, input dto2.CheckInput) (dto.CheckOutput, error) {
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

func (s *Usecase) CheckRole(ctx context.Context, input dto2.CheckRoleInput) (dto.CheckRoleOutput, error) {
	res, err := s.client.CheckRole(ctx, &pb.CheckRoleRequest{
		ID: input.ID,
	})
	if err != nil {
		response := dto.NewCheckRoleOutput(int(res.GetRole()), err.Error(), 404)
		return response, err
	}
	response := dto.NewCheckRoleOutput(int(res.GetRole()), res.GetError(), 200)
	return response, nil
}
