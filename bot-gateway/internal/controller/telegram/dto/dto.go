package dto

import (
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
