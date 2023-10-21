package books_usecase

import (
	"context"
	"gateway/internal/controller/telegram/dto"
	books_dto "gateway/internal/domain/books/dto"
	"gateway/internal/domain/books/model"
	pb "github.com/estoniec/libraryProject/contracts/gen/go/books"
	"strconv"
)

type Service interface {
	Create(ctx context.Context, dto books_dto.CreateSearchDTO) error
	Find(ctx context.Context, dto books_dto.FindSearchDTO) (books_dto.FindSearchOutput, error)
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
	stringID := strconv.Itoa(int(input.ID))
	newDTO := books_dto.NewCreateDTO(stringID, searched)
	err := u.bookService.Create(ctx, newDTO)
	if err != nil {
		response := books_dto.NewCreateOutput(err.Error())
		return response, err
	}
	return books_dto.NewCreateOutput(""), nil
}

func (u *Usecase) FindSearch(ctx context.Context, input dto.FindSearchInput) (books_dto.FindSearchOutput, error) {
	stringID := strconv.Itoa(int(input.ID))
	newDTO := books_dto.NewFindDTO(stringID)
	res, err := u.bookService.Find(ctx, newDTO)
	if err != nil {
		response := books_dto.NewFindOutput(err.Error(), nil)
		return response, err
	}
	return books_dto.NewFindOutput("", res.Searched), nil
}

func (u *Usecase) AddBook(ctx context.Context, input dto.AddBookInput) (books_dto.AddBookOutput, error) {
	dto := books_dto.NewAddBookDTO(input.Book)
	res, err := u.client.CreateBook(ctx, &pb.CreateBookRequest{
		Book: dto.Book,
	})
	if err != nil {
		response := books_dto.NewAddBookOutput(err.Error(), 404, input.Book)
		return response, err
	}
	response := books_dto.NewAddBookOutput(res.GetError(), res.GetStatus(), input.Book)
	return response, nil
}

func (u *Usecase) EditCountBook(ctx context.Context, input dto.EditCountBookInput) (books_dto.EditCountBookOutput, error) {
	dto := books_dto.NewEditCountBookDTO(input.ISBN, int64(input.Count))
	res, err := u.client.EditCountBook(ctx, &pb.EditCountBookRequest{
		ISBN:  dto.ISBN,
		Count: dto.Count,
	})
	if err != nil {
		response := books_dto.NewEditCountBookOutput(err.Error(), 404)
		return response, err
	}
	response := books_dto.NewEditCountBookOutput(res.GetError(), res.GetStatus())
	return response, nil
}
