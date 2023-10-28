package dto

import "rent/internal/domain/rent/model"

type RentBookInput struct {
	BookID   int64
	UserID   int64
	ReturnAt int64
}

type RentBookOutput struct {
	Error  string
	Status int64
	ID     int64
}

type CreateDTO struct {
	BookID   int
	UserID   int
	ReturnAt int
}

func NewRentBookOutput(error string, status int64, id int64) RentBookOutput {
	return RentBookOutput{
		Error:  error,
		Status: status,
		ID:     id,
	}
}

func NewCreateDTO(bookID int, userID int, returnAt int) CreateDTO {
	return CreateDTO{
		BookID:   bookID,
		UserID:   userID,
		ReturnAt: returnAt,
	}
}

type FindBookInput struct {
	ID int64
}

type FindBookOutput struct {
	Error  string
	Status int64
	Book   model.Book
}

func NewFindBookOutput(error string, status int64, book model.Book) FindBookOutput {
	return FindBookOutput{
		Error:  error,
		Status: status,
		Book:   book,
	}
}

type ConfirmRentInput struct {
	ID int64
}

type ConfirmRentOutput struct {
	Error  string
	Status int64
}

func NewConfirmRentOutput(error string, status int64) ConfirmRentOutput {
	return ConfirmRentOutput{
		Error:  error,
		Status: status,
	}
}

type GetDebtInput struct {
	Time int64
}

type GetDebtOutput struct {
	Error  string
	Status int64
	Debt   []model.BooksUsers
}

func NewGetDebtOutput(error string, status int64, debt []model.BooksUsers) GetDebtOutput {
	return GetDebtOutput{
		Error:  error,
		Status: status,
		Debt:   debt,
	}
}
