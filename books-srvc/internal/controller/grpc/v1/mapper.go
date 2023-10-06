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
	book := &pb.Book{
		ID:     int64(output.Book.ID),
		ISBN:   output.Book.ISBN,
		Name:   output.Book.Name,
		Author: output.Book.Author,
		Count:  int64(output.Book.Count),
	}
	return &pb.FindByISBNResponse{
		Status: output.Status,
		Error:  output.Error,
		Book:   book,
	}
}

func NewFindByAuthorInput(req *pb.FindByAuthorRequest) dto.FindByAuthorInput {
	return dto.FindByAuthorInput{
		Author: req.GetAuthor(),
		Offset: int(req.GetOffset()),
	}
}

func NewFindByAuthorOutput(output dto.FindByAuthorOutput) *pb.FindByAuthorResponse {
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
	return &pb.FindByAuthorResponse{
		Status: output.Status,
		Error:  output.Error,
		Books:  books,
	}
}

func NewFindByNameInput(req *pb.FindByNameRequest) dto.FindByNameInput {
	return dto.FindByNameInput{
		Name:   req.GetName(),
		Offset: int(req.GetOffset()),
	}
}

func NewFindByNameOutput(output dto.FindByNameOutput) *pb.FindByNameResponse {
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
	return &pb.FindByNameResponse{
		Status: output.Status,
		Error:  output.Error,
		Books:  books,
	}
}

func NewFindByNameAndAuthorInput(req *pb.FindByNameAndAuthorRequest) dto.FindByNameAndAuthorInput {
	return dto.FindByNameAndAuthorInput{
		Name:   req.GetName(),
		Author: req.GetAuthor(),
	}
}

func NewFindByNameAndAuthorOutput(output dto.FindByNameAndAuthorOutput) *pb.FindByNameAndAuthorResponse {
	book := &pb.Book{
		ID:     int64(output.Book.ID),
		ISBN:   output.Book.ISBN,
		Name:   output.Book.Name,
		Author: output.Book.Author,
		Count:  int64(output.Book.Count),
	}
	return &pb.FindByNameAndAuthorResponse{
		Status: output.Status,
		Error:  output.Error,
		Books:  book,
	}
}

func NewFindAllInput(req *pb.FindAllRequest) dto.FindAllInput {
	return dto.FindAllInput{
		Offset: int(req.GetOffset()),
	}
}

func NewFindAllOutput(output dto.FindAllOutput) *pb.FindAllResponse {
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
	return &pb.FindAllResponse{
		Status: output.Status,
		Error:  output.Error,
		Books:  books,
	}
}
