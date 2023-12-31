package booksService

import (
	"gateway/internal/domain/books/model"
	pb "github.com/estoniec/libraryProject/contracts/gen/go/books"
)

type FindByOutput struct {
	Error  string
	Status int64
	Book   []model.Book
}

type CreateSearchOutput struct {
	Error string
}

type CreateSearchDTO struct {
	ID       string   `json:"id"`
	Searched []string `json:"searched"`
}

type FindSearchOutput struct {
	Error    string
	ID       string   `json:"id"`
	Searched []string `json:"searched"`
}

type FindSearchDTO struct {
	ID string
}

type AddBookOutput struct {
	Error  string
	Status int64
	Book   model.Book
}

type AddBookDTO struct {
	Book *pb.Book
}

type EditCountBookOutput struct {
	Error  string
	Status int64
}

type EditCountBookDTO struct {
	ISBN  string
	Count int64
}

type DeleteBookOutput struct {
	Error  string
	Status int64
}

func NewByOutput(error string, status int64, books []model.Book) FindByOutput {
	return FindByOutput{
		Error:  error,
		Status: status,
		Book:   books,
	}
}

func NewCreateOutput(error string) CreateSearchOutput {
	return CreateSearchOutput{
		Error: error,
	}
}

func NewCreateDTO(ID string, searched []string) CreateSearchDTO {
	return CreateSearchDTO{
		ID:       ID,
		Searched: searched,
	}
}

func NewFindOutput(error string, searched []string) FindSearchOutput {
	return FindSearchOutput{
		Error:    error,
		Searched: searched,
	}
}

func NewFindDTO(ID string) FindSearchDTO {
	return FindSearchDTO{
		ID: ID,
	}
}

func NewAddBookOutput(error string, status int64, book model.Book) AddBookOutput {
	return AddBookOutput{
		Error:  error,
		Status: status,
		Book:   book,
	}
}

func NewAddBookDTO(book model.Book) AddBookDTO {
	return AddBookDTO{
		Book: &pb.Book{
			ISBN:   book.ISBN,
			Count:  int64(book.Count),
			Name:   book.Name,
			Author: book.Author,
		},
	}
}

func NewEditCountBookOutput(error string, status int64) EditCountBookOutput {
	return EditCountBookOutput{
		Error:  error,
		Status: status,
	}
}

func NewEditCountBookDTO(isbn string, count int64) EditCountBookDTO {
	return EditCountBookDTO{
		ISBN:  isbn,
		Count: count,
	}
}

func NewDeleteBookOutput(error string, status int64) DeleteBookOutput {
	return DeleteBookOutput{
		Error:  error,
		Status: status,
	}
}
