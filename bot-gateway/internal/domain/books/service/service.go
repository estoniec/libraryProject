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
	book := model.NewBook(int(res.GetId()), res.GetIsbn(), int(res.GetCount()), res.GetName(), res.GetAuthor())
	if err != nil {
		response := dto.NewISBNOutput(err.Error(), 404, book)
		return response, err
	}
	response := dto.NewISBNOutput(res.GetError(), res.GetStatus(), book)
	return response, nil
}
