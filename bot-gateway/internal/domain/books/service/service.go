package books_service

import (
	"context"
	dto "gateway/internal/domain/books/dto"
	"gateway/internal/domain/books/model"
	pb "github.com/estoniec/libraryProject/contracts/gen/go/books"
)

type Service struct {
	client pb.BooksServiceClient
}

func NewService(client pb.BooksServiceClient) *Service {
	return &Service{
		client: client,
	}
}

func (s *Service) FindByISBN(ctx context.Context, input dto.FindByISBNInput) (dto.FindByISBNOutput, error) {
	res, err := s.client.FindByISBN(ctx, &pb.FindByISBNRequest{
		ISBN: input.ISBN,
	})
	book := model.NewBook(res.GetBook())
	if err != nil {
		response := dto.NewISBNOutput(err.Error(), 404, book)
		return response, err
	}
	response := dto.NewISBNOutput(res.GetError(), res.GetStatus(), book)
	return response, nil
}

func (s *Service) FindByAuthor(ctx context.Context, input dto.FindByAuthorInput) (dto.FindByAuthorOutput, error) {
	res, err := s.client.FindByAuthor(ctx, &pb.FindByAuthorRequest{
		Author: input.Author,
		Offset: input.Offset,
	})
	books := model.NewBooks(res.GetBooks())
	if err != nil {
		response := dto.NewAuthorOutput(err.Error(), 404, books)
		return response, err
	}
	response := dto.NewAuthorOutput(res.GetError(), res.GetStatus(), books)
	return response, nil
}

func (s *Service) FindByName(ctx context.Context, input dto.FindByNameInput) (dto.FindByNameOutput, error) {
	res, err := s.client.FindByName(ctx, &pb.FindByNameRequest{
		Name:   input.Name,
		Offset: input.Offset,
	})
	books := model.NewBooks(res.GetBooks())
	if err != nil {
		response := dto.NewNameOutput(err.Error(), 404, books)
		return response, err
	}
	response := dto.NewNameOutput(res.GetError(), res.GetStatus(), books)
	return response, nil
}
