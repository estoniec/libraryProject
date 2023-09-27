package v1

import (
	"books-srvc/internal/domain/books/dto"
	pb "github.com/estoniec/libraryProject/contracts/gen/go/books"
)

func NewFindByISBNInput(req *pb.FindByISBNRequest) dto.FindByISBNInput {
	return dto.FindByISBNInput{
		ISBN: req.GetISBN(),
	}
}

func NewFindByISBNOutput(output dto.FindByISBNOutput) *pb.FindByISBNResponse {
	return &pb.FindByISBNResponse{
		Status: output.Status,
		Error:  output.Error,
		Isbn:   output.Book.ISBN,
		Id:     int64(output.Book.ID),
		Count:  int64(output.Book.Count),
		Author: output.Book.Author,
		Name:   output.Book.Name,
	}
}
