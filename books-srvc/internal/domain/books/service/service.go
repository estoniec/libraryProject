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
	FindByNameAndAuthor(ctx context.Context, dto dto.FindByNameAndAuthorInput) (model.Book, error)
	FindAll(ctx context.Context, dto dto.FindAllInput) ([]model.Book, error)
	FindBy(ctx context.Context, dto dto.FindByInput) ([]model.Book, error)
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

func (s *Service) FindByNameAndAuthor(ctx context.Context, input dto.FindByNameAndAuthorInput) (dto.FindByNameAndAuthorOutput, error) {
	book, err := s.repository.FindByNameAndAuthor(ctx, input)
	if err != nil {
		return dto.NewFindByNameAndAuthorOutput(err.Error(), 404, book), err
	}
	return dto.NewFindByNameAndAuthorOutput("", 200, book), nil
}

func (s *Service) FindAll(ctx context.Context, input dto.FindAllInput) (dto.FindAllOutput, error) {
	books, err := s.repository.FindAll(ctx, input)
	if err != nil {
		return dto.NewFindAllOutput(err.Error(), 404, books), err
	}
	return dto.NewFindAllOutput("", 200, books), nil
}

func (s *Service) FindBy(ctx context.Context, input dto.FindByInput) (dto.FindByOutput, error) {
	books, err := s.repository.FindBy(ctx, input)
	if err != nil {
		return dto.NewFindByOutput(err.Error(), 404, books), err
	}
	return dto.NewFindByOutput("", 200, books), nil
}
