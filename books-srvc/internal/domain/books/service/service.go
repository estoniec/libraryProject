package service

import (
	"books-srvc/internal/domain/books/dto"
	"books-srvc/internal/domain/books/model"
	"context"
)

type Repository interface {
	FindBy(ctx context.Context, dto dto.FindByInput) ([]model.Book, error)
	Create(ctx context.Context, dto dto.CreateBookInput) (model.Book, error)
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) FindBy(ctx context.Context, input dto.FindByInput) (dto.FindByOutput, error) {
	books, err := s.repository.FindBy(ctx, input)
	if err != nil {
		return dto.NewFindByOutput(err.Error(), 404, books), err
	}
	return dto.NewFindByOutput("", 200, books), nil
}

func (s *Service) CreateBook(ctx context.Context, input dto.CreateBookInput) (dto.CreateBookOutput, error) {
	books, err := s.repository.Create(ctx, input)
	if err != nil {
		return dto.NewCreateBookOutput(err.Error(), 404, books), err
	}
	return dto.NewCreateBookOutput("", 200, books), nil
}
