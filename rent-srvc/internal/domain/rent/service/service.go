package service

import (
	"context"
	"rent/internal/domain/rent/dto"
)

type Repository interface {
	Create(ctx context.Context, dto dto.CreateDTO) (int64, error)
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
		return dto.NewRentBookOutput(err.Error(), 404, id), err
	}
	return dto.NewRentBookOutput(err.Error(), 404, id), nil
}
