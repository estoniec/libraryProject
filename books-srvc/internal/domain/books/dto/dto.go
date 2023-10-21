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

type CreateBookInput struct {
	Book model.Book
}

type CreateBookOutput struct {
	Error  string
	Status int64
	Book   model.Book
}

type EditCountBookInput struct {
	ISBN  string
	Count int
}

type EditCountBookOutput struct {
	Error  string
	Status int64
}

type DeleteBookInput struct {
	ISBN string
}

type DeleteBookOutput struct {
	Error  string
	Status int64
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

func NewCreateBookOutput(error string,
	status int64,
	book model.Book) CreateBookOutput {
	return CreateBookOutput{
		Error:  error,
		Status: status,
		Book:   book,
	}
}

func NewEditCountOutput(error string,
	status int64) EditCountBookOutput {
	return EditCountBookOutput{
		Error:  error,
		Status: status,
	}
}

func NewDeleteOutput(error string,
	status int64) DeleteBookOutput {
	return DeleteBookOutput{
		Error:  error,
		Status: status,
	}
}
