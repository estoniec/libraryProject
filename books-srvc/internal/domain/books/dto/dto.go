package dto

import "books-srvc/internal/domain/books/model"

type FindByInput struct {
	Offset int
	Book   model.Book
}

type FindByOutput struct {
	Error  string
	Status int64
	Book   []model.Book
}

func NewFindByOutput(error string,
	status int64,
	book []model.Book) FindByOutput {
	return FindByOutput{
		Error:  error,
		Status: status,
		Book:   book,
	}
}
