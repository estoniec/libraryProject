package service

import (
	"context"
	"rent/internal/domain/rent/dto"
)

type Repository interface {
	Create(ctx context.Context, dto dto.CreateDTO) (string, string, error)
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
	author, name, err := s.repository.Create(ctx, dtoRent)
	if err != nil {
		return dto.NewRentBookOutput(author, name, err.Error(), 404), err
	}
	return dto.NewRentBookOutput(author, name, err.Error(), 404), nil
}
