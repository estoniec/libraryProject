package service

import (
	"books-srvc/internal/domain/books/dto"
	"books-srvc/internal/domain/books/model"
	"context"
)

type Repository interface {
	FindByISBN(ctx context.Context, dto dto.FindByISBNInput) (model.Book, error)
	FindByAuthor(ctx context.Context, dto dto.FindByAuthorInput) ([]model.Book, error)
	FindByName(ctx context.Context, dto dto.FindByNameInput) ([]model.Book, error)
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

func (s *Service) FindByAuthor(ctx context.Context, input dto.FindByAuthorInput) (dto.FindByAuthorOutput, error) {
	book, err := s.repository.FindByAuthor(ctx, input)
	if err != nil {
		return dto.NewFindByAuthorOutput(err.Error(), 404, book), err
	}
	return dto.NewFindByAuthorOutput("", 200, book), nil
}

func (s *Service) FindByName(ctx context.Context, input dto.FindByNameInput) (dto.FindByNameOutput, error) {
	book, err := s.repository.FindByName(ctx, input)
	if err != nil {
		return dto.NewFindByNameOutput(err.Error(), 404, book), err
	}
	return dto.NewFindByNameOutput("", 200, book), nil
}
