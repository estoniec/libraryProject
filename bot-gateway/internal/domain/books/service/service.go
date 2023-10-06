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

func (s *Service) FindByNameAndAuthor(ctx context.Context, input dto.FindByNameAndAuthorInput) (dto.FindByNameAndAuthorOutput, error) {
	res, err := s.client.FindByNameAndAuthor(ctx, &pb.FindByNameAndAuthorRequest{
		Author: input.Author,
		Name:   input.Name,
	})
	books := model.NewBook(res.GetBooks())
	if err != nil {
		response := dto.NameAndAuthorOutput(err.Error(), 404, books)
		return response, err
	}
	response := dto.NameAndAuthorOutput(res.GetError(), res.GetStatus(), books)
	return response, nil
}

func (s *Service) FindAll(ctx context.Context, input dto.FindAllInput) (dto.FindAllOutput, error) {
	res, err := s.client.FindAll(ctx, &pb.FindAllRequest{
		Offset: input.Offset,
	})
	books := model.NewBooks(res.GetBooks())
	if err != nil {
		response := dto.NewAllOutput(err.Error(), 404, books)
		return response, err
	}
	response := dto.NewAllOutput(res.GetError(), res.GetStatus(), books)
	return response, nil
}
