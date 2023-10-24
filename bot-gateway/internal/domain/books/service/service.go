package booksService

import (
	"context"
	books_dto "gateway/internal/domain/books/dto"
	pb "github.com/estoniec/libraryProject/contracts/gen/go/books"
)

type Storage interface {
	Create(ctx context.Context, dto books_dto.CreateSearchDTO) error
	Find(ctx context.Context, dto books_dto.FindSearchDTO) (books_dto.FindSearchOutput, error)
}

type Service struct {
	client  pb.BooksServiceClient
	storage Storage
}

func NewService(client pb.BooksServiceClient, storage Storage) *Service {
	return &Service{
		client:  client,
		storage: storage,
	}
}

func (s *Service) Create(ctx context.Context, dto books_dto.CreateSearchDTO) error {
	return s.storage.Create(ctx, dto)
}

func (s *Service) Find(ctx context.Context, dto books_dto.FindSearchDTO) (books_dto.FindSearchOutput, error) {
	return s.storage.Find(ctx, dto)
}
