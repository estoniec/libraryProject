package dto

import (
	"gateway/internal/domain/books/model"
	pb "github.com/estoniec/libraryProject/contracts/gen/go/books"
)

type FindByInput struct {
	Offset int
	Book   *pb.FindBook
}

type CreateSearchInput struct {
	ID        int64
	FindType  string
	Something string
}

type CheckRoleInput struct {
	ID int64
}

func NewByInput(offset int64, book *pb.FindBook) FindByInput {
	return FindByInput{
		Offset: int(offset),
		Book:   book,
	}
}

func NewCreateInput(ID int64, findType string, something string) CreateSearchInput {
	return CreateSearchInput{
		ID:        ID,
		FindType:  findType,
		Something: something,
	}
}

type FindSearchInput struct {
	ID int64
}

func NewFindInput(ID int64) FindSearchInput {
	return FindSearchInput{
		ID: ID,
	}
}

func NewCheckRoleInput(ID int64) CheckRoleInput {
	return CheckRoleInput{
		ID: ID,
	}
}

type AddBookInput struct {
	Book model.Book
}

func NewAddBookInput(book model.Book) AddBookInput {
	return AddBookInput{
		Book: book,
	}
}

type RegInput struct {
	Phone    string
	Username string
	Class    string
	ID       int64
}

type CheckInput struct {
	ID int64
}

func NewRegInput(phone string, username string, class string, ID int64) RegInput {
	return RegInput{
		Phone:    phone,
		Username: username,
		Class:    class,
		ID:       ID,
	}
}

func NewCheckInput(id int64) CheckInput {
	return CheckInput{
		ID: id,
	}
}

type EditCountBookInput struct {
	ISBN  string
	Count int
}

func NewEditCountBookInput(isbn string, count int) EditCountBookInput {
	return EditCountBookInput{
		ISBN:  isbn,
		Count: count,
	}
}

type DeleteBookInput struct {
	ISBN string
}

func NewDeleteBookInput(isbn string) DeleteBookInput {
	return DeleteBookInput{
		ISBN: isbn,
	}
}

type RentInput struct {
	BookID   string
	UserID   int64
	ReturnAt int64
}

func NewRentInput(bookID string, userID int64, returnAt int64) RentInput {
	return RentInput{
		BookID:   bookID,
		UserID:   userID,
		ReturnAt: returnAt,
	}
}

type FindBookInput struct {
	ID int64
}

func NewFindBookInput(id int64) FindBookInput {
	return FindBookInput{
		ID: id,
	}
}

type ConfirmRentInput struct {
	ID int64
}

func NewConfirmRentInput(id int64) ConfirmRentInput {
	return ConfirmRentInput{
		ID: id,
	}
}

type GetDebtInput struct {
	Time   int64
	Offset int64
}

func NewGetDebtInput(time int64, offset int64) GetDebtInput {
	return GetDebtInput{
		Time:   time,
		Offset: offset,
	}
}

type CheckRentInput struct {
	UID int64
	BID int64
}

func NewCheckRentInput(uid int64, bid int64) CheckRentInput {
	return CheckRentInput{
		UID: uid,
		BID: bid,
	}
}

type ConfirmReturnInput struct {
	ID int64
}

func NewConfirmReturnInput(id int64) ConfirmReturnInput {
	return ConfirmReturnInput{
		ID: id,
	}
}

type MyRentsInput struct {
	UserID int64
	Offset int64
}

func NewMyRentsInput(userid int64, offset int64) MyRentsInput {
	return MyRentsInput{
		UserID: userid,
		Offset: offset,
	}
}
