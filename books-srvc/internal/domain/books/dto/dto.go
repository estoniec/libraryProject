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

type FindByAuthorInput struct {
	Author string
	Offset int
}

type FindByAuthorOutput struct {
	Error  string
	Status int64
	Book   []model.Book
}

type FindByNameInput struct {
	Name   string
	Offset int
}

type FindByNameOutput struct {
	Error  string
	Status int64
	Book   []model.Book
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

func NewFindByAuthorOutput(error string,
	status int64,
	book []model.Book) FindByAuthorOutput {
	return FindByAuthorOutput{
		Error:  error,
		Status: status,
		Book:   book,
	}
}

func NewFindByNameOutput(error string,
	status int64,
	book []model.Book) FindByNameOutput {
	return FindByNameOutput{
		Error:  error,
		Status: status,
		Book:   book,
	}
}
