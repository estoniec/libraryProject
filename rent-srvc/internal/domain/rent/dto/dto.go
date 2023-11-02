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

type ConfirmReturnInput struct {
	ID int64
}

type ConfirmReturnOutput struct {
	Error  string
	Status int64
}

func NewConfirmReturnOutput(error string, status int64) ConfirmReturnOutput {
	return ConfirmReturnOutput{
		Error:  error,
		Status: status,
	}
}

type FindByInput struct {
	Offset int64
	ID     int64
	Time   int64
	UserID int64
	BookID int64
}

type FindByOutput struct {
	Error  string
	Status int64
	Model  []model.BooksUsers
}

func NewFindByOutput(error string, status int64, model []model.BooksUsers) FindByOutput {
	return FindByOutput{
		Error:  error,
		Status: status,
		Model:  model,
	}
}
