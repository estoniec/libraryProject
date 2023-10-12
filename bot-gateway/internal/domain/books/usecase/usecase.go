package books_usecase

import (
	"context"
	dto "gateway/internal/domain/books/dto"
	"gateway/internal/domain/books/model"
	books_service "gateway/internal/domain/books/service"
	pb "github.com/estoniec/libraryProject/contracts/gen/go/books"
)

type Usecase struct {
	client      pb.BooksServiceClient
	bookService books_service.Service
}

func NewUsecase(client pb.BooksServiceClient, booksService books_service.Service) *Usecase {
	return &Usecase{
		client:      client,
		bookService: booksService,
	}
}

func (s *Usecase) FindBy(ctx context.Context, input dto.FindByInput) (dto.FindByOutput, error) {
	res, err := s.client.FindBy(ctx, &pb.FindByRequest{
		Offset: int64(input.Offset),
		Find:   input.Book,
	})
	books := model.NewBooks(res.GetBook())
	if err != nil {
		response := dto.NewByOutput(err.Error(), 404, books)
		return response, err
	}
	response := dto.NewByOutput(res.GetError(), res.GetStatus(), books)
	return response, nil
}

func (s *Usecase) CreateSearch(ctx context.Context, input dto.FindByInput) (dto.FindByOutput, error) {
	res, err := s.client.FindBy(ctx, &pb.FindByRequest{
		Offset: int64(input.Offset),
		Find:   input.Book,
	})
	books := model.NewBooks(res.GetBook())
	if err != nil {
		response := dto.NewByOutput(err.Error(), 404, books)
		return response, err
	}
	response := dto.NewByOutput(res.GetError(), res.GetStatus(), books)
	return response, nil
}

func (s *Usecase) FindSearch(ctx context.Context, input dto.FindByInput) (dto.FindByOutput, error) {
	res, err := s.client.FindBy(ctx, &pb.FindByRequest{
		Offset: int64(input.Offset),
		Find:   input.Book,
	})
	books := model.NewBooks(res.GetBook())
	if err != nil {
		response := dto.NewByOutput(err.Error(), 404, books)
		return response, err
	}
	response := dto.NewByOutput(res.GetError(), res.GetStatus(), books)
	return response, nil
}

func (s *Usecase) DeleteSearch(ctx context.Context, input dto.FindByInput) (dto.FindByOutput, error) {
	res, err := s.client.FindBy(ctx, &pb.FindByRequest{
		Offset: int64(input.Offset),
		Find:   input.Book,
	})
	books := model.NewBooks(res.GetBook())
	if err != nil {
		response := dto.NewByOutput(err.Error(), 404, books)
		return response, err
	}
	response := dto.NewByOutput(res.GetError(), res.GetStatus(), books)
	return response, nil
}
