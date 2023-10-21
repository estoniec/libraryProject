package service

import (
	"books-srvc/internal/domain/books/dto"
	"books-srvc/internal/domain/books/model"
	"context"
)

type Repository interface {
	FindBy(ctx context.Context, dto dto.FindByInput) ([]model.Book, error)
	Create(ctx context.Context, dto dto.CreateBookInput) (model.Book, error)
	EditCount(ctx context.Context, dto dto.EditCountBookInput) error
	Delete(ctx context.Context, dto dto.DeleteBookInput) error
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

func (s *Service) EditCountBook(ctx context.Context, input dto.EditCountBookInput) (dto.EditCountBookOutput, error) {
	err := s.repository.EditCount(ctx, input)
	if err != nil {
		return dto.NewEditCountOutput(err.Error(), 404), err
	}
	return dto.NewEditCountOutput("", 200), nil
}

func (s *Service) DeleteBook(ctx context.Context, input dto.DeleteBookInput) (dto.DeleteBookOutput, error) {
	err := s.repository.Delete(ctx, input)
	if err != nil {
		return dto.NewDeleteOutput(err.Error(), 404), err
	}
	return dto.NewDeleteOutput("", 200), nil
}
