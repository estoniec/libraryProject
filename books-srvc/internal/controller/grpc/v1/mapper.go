package v1

import (
	"books-srvc/internal/domain/books/dto"
	"books-srvc/internal/domain/books/model"
	pb "github.com/estoniec/libraryProject/contracts/gen/go/books"
)

func NewFindByInput(req *pb.FindByRequest) dto.FindByInput {
	return dto.FindByInput{
		Offset: int(req.GetOffset()),
		Book:   model.NewFindBook(req.GetFind().GetISBN(), req.GetFind().GetName(), req.GetFind().GetAuthor()),
	}
}

func NewFindByOutput(output dto.FindByOutput) *pb.FindByResponse {
	var books []*pb.Book
	if len(output.Book) == 0 {
		books = nil
	} else {
		books = make([]*pb.Book, len(output.Book))
		for i, book := range output.Book {
			books[i] = &pb.Book{
				ID:     int64(book.ID),
				ISBN:   book.ISBN,
				Name:   book.Name,
				Author: book.Author,
				Count:  int64(book.Count),
			}
		}
	}
	return &pb.FindByResponse{
		Status: output.Status,
		Error:  output.Error,
		Book:   books,
	}
}

func NewCreateInput(req *pb.CreateBookRequest) dto.CreateBookInput {
	return dto.CreateBookInput{
		Book: model.NewBook(int(req.GetBook().GetCount()), req.GetBook().GetISBN(), req.GetBook().GetName(), req.GetBook().GetAuthor()),
	}
}

func NewCreateOutput(output dto.CreateBookOutput) *pb.CreateBookResponse {
	return &pb.CreateBookResponse{
		Status: output.Status,
		Error:  output.Error,
		Books: &pb.Book{
			ID:     int64(output.Book.ID),
			ISBN:   output.Book.ISBN,
			Name:   output.Book.Name,
			Author: output.Book.Author,
			Count:  int64(output.Book.Count),
		},
	}
}
