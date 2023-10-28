package rentService

import (
	"gateway/internal/domain/rent/model"
	pb "github.com/estoniec/libraryProject/contracts/gen/go/books_users"
)

type RentBookOutput struct {
	Error  string
	Status int64
	ID     int64
}

type RentBookDTO struct {
	BookID   int64
	UserID   int64
	ReturnAt int64
}

func NewRentBookOutput(error string, status int64, id int64) RentBookOutput {
	return RentBookOutput{
		Error:  error,
		Status: status,
		ID:     id,
	}
}

func NewRentBookDTO(bookID int64, userID int64, returnAt int64) RentBookDTO {
	return RentBookDTO{
		BookID:   bookID,
		UserID:   userID,
		ReturnAt: returnAt,
	}
}

type FindBookOutput struct {
	Error  string
	Status int64
	Book   *pb.Book
}

func NewFindBookOutput(error string, status int64, book *pb.Book) FindBookOutput {
	return FindBookOutput{
		Error:  error,
		Status: status,
		Book:   book,
	}
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

type GetDebtOutput struct {
	Error  string
	Status int64
	Debt   []rentService.BooksUsers
}

func NewGetDebtOutput(error string, status int64, debts []rentService.BooksUsers) GetDebtOutput {
	return GetDebtOutput{
		Error:  error,
		Status: status,
		Debt:   debts,
	}
}
