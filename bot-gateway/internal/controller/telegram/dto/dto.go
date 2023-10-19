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
