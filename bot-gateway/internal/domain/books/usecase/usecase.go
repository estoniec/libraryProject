package books_usecase

import (
	"context"
	"gateway/internal/controller/telegram/dto"
	books_dto "gateway/internal/domain/books/dto"
	"gateway/internal/domain/books/model"
	pb "github.com/estoniec/libraryProject/contracts/gen/go/books"
)

type Service interface {
	Create(ctx context.Context, dto books_dto.CreateSearchDTO) error
}

type Usecase struct {
	client      pb.BooksServiceClient
	bookService Service
}

func NewUsecase(client pb.BooksServiceClient, booksService Service) *Usecase {
	return &Usecase{
		client:      client,
		bookService: booksService,
	}
}

func (u *Usecase) FindBy(ctx context.Context, input dto.FindByInput) (books_dto.FindByOutput, error) {
	res, err := u.client.FindBy(ctx, &pb.FindByRequest{
		Offset: int64(input.Offset),
		Find:   input.Book,
	})
	books := model.NewBooks(res.GetBook())
	if err != nil {
		response := books_dto.NewByOutput(err.Error(), 404, books)
		return response, err
	}
	response := books_dto.NewByOutput(res.GetError(), res.GetStatus(), books)
	return response, nil
}

func (u *Usecase) CreateSearch(ctx context.Context, input dto.CreateSearchInput) (books_dto.CreateSearchOutput, error) {
	var searched []string
	searched = append(searched, input.FindType)
	searched = append(searched, input.Something)
	newDTO := books_dto.NewCreateDTO(string(input.ID), searched)
	err := u.bookService.Create(ctx, newDTO)
	if err != nil {
		response := books_dto.NewCreateOutput(err.Error())
		return response, err
	}
	return books_dto.NewCreateOutput(""), nil
	//res, err := s.client.FindBy(ctx, &pb.FindByRequest{
	//	Offset: int64(input.Offset),
	//	Find:   input.Book,
	//})
	//books := model.NewBooks(res.GetBook())
	//if err != nil {
	//	response := dto.NewByOutput(err.Error(), 404, books)
	//	return response, err
	//}
	//response := dto.NewByOutput(res.GetError(), res.GetStatus(), books)
	//return response, nil
}

//func (u *Usecase) FindSearch(ctx context.Context, input dto.FindByInput) (books_dto.FindByOutput, error) {
//	res, err := u.client.FindBy(ctx, &pb.FindByRequest{
//		Offset: int64(input.Offset),
//		Find:   input.Book,
//	})
//	books := model.NewBooks(res.GetBook())
//	if err != nil {
//		response := dto.NewByOutput(err.Error(), 404, books)
//		return response, err
//	}
//	response := dto.NewByOutput(res.GetError(), res.GetStatus(), books)
//	return response, nil
//}
//
//func (u *Usecase) DeleteSearch(ctx context.Context, input dto.FindByInput) (books_dto.FindByOutput, error) {
//	res, err := u.client.FindBy(ctx, &pb.FindByRequest{
//		Offset: int64(input.Offset),
//		Find:   input.Book,
//	})
//	books := model.NewBooks(res.GetBook())
//	if err != nil {
//		response := dto.NewByOutput(err.Error(), 404, books)
//		return response, err
//	}
//	response := dto.NewByOutput(res.GetError(), res.GetStatus(), books)
//	return response, nil
//}
