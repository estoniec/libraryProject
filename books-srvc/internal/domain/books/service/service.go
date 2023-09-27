package service

import (
	"books-srvc/internal/domain/books/dto"
	"books-srvc/internal/domain/books/model"
	"context"
)

type Repository interface {
	FindByISBN(ctx context.Context, dto dto.FindByISBNInput) (model.Book, error)
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) FindByISBN(ctx context.Context, input dto.FindByISBNInput) (dto.FindByISBNOutput, error) {
	book, err := s.repository.FindByISBN(ctx, input)
	if err != nil {
		return dto.NewFindByISBNOutput(err.Error(), 404, book), err
	}
	return dto.NewFindByISBNOutput("", 200, book), nil
}
