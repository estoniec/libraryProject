package service

import (
	"context"
	"rent/internal/domain/rent/dto"
	"rent/internal/domain/rent/model"
)

type Repository interface {
	Create(ctx context.Context, dto dto.CreateDTO) (int, error)
	UpdateGet(ctx context.Context, dto dto.ConfirmRentInput) error
	UpdateReturn(ctx context.Context, dto dto.ConfirmReturnInput) error
	FindBy(ctx context.Context, dto dto.FindByInput) ([]model.BooksUsers, error)
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) RentBook(ctx context.Context, input dto.RentBookInput) (dto.RentBookOutput, error) {
	dtoRent := dto.NewCreateDTO(int(input.BookID), int(input.UserID), int(input.ReturnAt))
	id, err := s.repository.Create(ctx, dtoRent)
	if err != nil {
		return dto.NewRentBookOutput(err.Error(), 404, 0), err
	}
	return dto.NewRentBookOutput("", 200, int64(id)), nil
}

func (s *Service) ConfirmRent(ctx context.Context, input dto.ConfirmRentInput) (dto.ConfirmRentOutput, error) {
	err := s.repository.UpdateGet(ctx, input)
	if err != nil {
		return dto.NewConfirmRentOutput(err.Error(), 404), err
	}
	return dto.NewConfirmRentOutput("", 200), nil
}

func (s *Service) ConfirmReturn(ctx context.Context, input dto.ConfirmReturnInput) (dto.ConfirmReturnOutput, error) {
	err := s.repository.UpdateReturn(ctx, input)
	if err != nil {
		return dto.NewConfirmReturnOutput(err.Error(), 404), err
	}
	return dto.NewConfirmReturnOutput("", 200), nil
}

func (s *Service) FindBy(ctx context.Context, input dto.FindByInput) (dto.FindByOutput, error) {
	model, err := s.repository.FindBy(ctx, input)
	if err != nil {
		return dto.NewFindByOutput(err.Error(), 404, model), err
	}
	return dto.NewFindByOutput("", 200, model), nil
}
