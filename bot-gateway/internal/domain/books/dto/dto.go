package books_dto

import (
	"gateway/internal/domain/books/model"
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
