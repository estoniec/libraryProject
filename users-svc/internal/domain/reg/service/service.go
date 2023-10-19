package service

import (
	"context"
	"registration-svc/internal/domain/reg/dto"
	"registration-svc/internal/domain/reg/model"
)

type Repository interface {
	CreateUser(ctx context.Context, req model.User) error
	FindUser(ctx context.Context, req dto.CheckInput) (dto.CheckOutput, error)
	FindUserByRole(ctx context.Context, req dto.CheckRoleInput) (dto.CheckRoleOutput, error)
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) RegUser(ctx context.Context, model model.User) (dto.RegOutput, error) {
	err := s.repository.CreateUser(ctx, model)
	if err != nil {
		return dto.NewRegOutput(err.Error(), 404), err
	}
	return dto.NewRegOutput("", 200), nil
}

func (s *Service) CheckUser(ctx context.Context, input dto.CheckInput) (dto.CheckOutput, error) {
	checked, err := s.repository.FindUser(ctx, input)
	if err != nil {
		return checked, err
	}
	return checked, nil
}

func (s *Service) CheckRole(ctx context.Context, input dto.CheckRoleInput) (dto.CheckRoleOutput, error) {
	role, err := s.repository.FindUserByRole(ctx, input)
	if err != nil {
		return role, err
	}
	return role, nil
}
