package reg_dto

import "gateway/internal/domain/books/model"

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
	Offset int64
}

type FindByAuthorOutput struct {
	Error  string
	Status int64
	Book   []model.Book
}

type FindByNameInput struct {
	Name   string
	Offset int64
}

type FindByNameOutput struct {
	Error  string
	Status int64
	Book   []model.Book
}

func NewISBNInput(isbn string) FindByISBNInput {
	return FindByISBNInput{
		ISBN: isbn,
	}
}

func NewISBNOutput(error string, status int64, book model.Book) FindByISBNOutput {
	return FindByISBNOutput{
		Error:  error,
		Status: status,
		Book:   book,
	}
}

func NewAuthorInput(author string, offset int64) FindByAuthorInput {
	return FindByAuthorInput{
		Author: author,
		Offset: offset,
	}
}

func NewAuthorOutput(error string, status int64, book []model.Book) FindByAuthorOutput {
	return FindByAuthorOutput{
		Error:  error,
		Status: status,
		Book:   book,
	}
}

func NewNameInput(name string, offset int64) FindByNameInput {
	return FindByNameInput{
		Name:   name,
		Offset: offset,
	}
}

func NewNameOutput(error string, status int64, book []model.Book) FindByNameOutput {
	return FindByNameOutput{
		Error:  error,
		Status: status,
		Book:   book,
	}
}
