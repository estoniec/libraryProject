package dto

import "books-srvc/internal/domain/books/model"

type FindByISBNInput struct {
	ISBN string
}

type FindByISBNOutput struct {
	Error  string
	Status int64
	Book   model.Book
}

func NewFindByISBNOutput(error string,
	status int64,
	book model.Book) FindByISBNOutput {
	return FindByISBNOutput{
		Error:  error,
		Status: status,
		Book:   book,
	}
}
