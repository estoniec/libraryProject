package books_service

import (
	"context"
	"gateway/internal/domain/books/dto"
	pb "github.com/estoniec/libraryProject/contracts/gen/go/books"
)

type Repository interface {
	Create(ctx context.Context, dto books_dto.CreateSearchDTO) error
}

type Service struct {
	client     pb.BooksServiceClient
	repository Repository
}

func NewService(client pb.BooksServiceClient, repository Repository) *Service {
	return &Service{
		client:     client,
		repository: repository,
	}
}

func (s *Service) Create(ctx context.Context, dto books_dto.CreateSearchDTO) error {
	return s.repository.Create(ctx, dto)
}
