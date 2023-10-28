package service

import (
	"context"
	"rent/internal/domain/rent/dto"
	"rent/internal/domain/rent/model"
)

type Repository interface {
	Create(ctx context.Context, dto dto.CreateDTO) (int, error)
	Find(ctx context.Context, dto dto.FindBookInput) (model.Book, error)
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

func (s *Service) FindBook(ctx context.Context, input dto.FindBookInput) (dto.FindBookOutput, error) {
	book, err := s.repository.Find(ctx, input)
	if err != nil {
		return dto.NewFindBookOutput(err.Error(), 404, book), err
	}
	return dto.NewFindBookOutput("", 200, book), nil
}
