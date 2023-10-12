package dto

import (
	"gateway/internal/domain/books/model"
	pb "github.com/estoniec/libraryProject/contracts/gen/go/books"
)

type FindByInput struct {
	Offset int
	Book   *pb.FindBook
}

type FindByOutput struct {
	Error  string
	Status int64
	Book   []model.Book
}

func NewByInput(offset int64, book *pb.FindBook) FindByInput {
	return FindByInput{
		Offset: int(offset),
		Book:   book,
	}
}

func NewByOutput(error string, status int64, books []model.Book) FindByOutput {
	return FindByOutput{
		Error:  error,
		Status: status,
		Book:   books,
	}
}
