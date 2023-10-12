package model

import pb "github.com/estoniec/libraryProject/contracts/gen/go/books"

type Book struct {
	ID     int
	ISBN   string
	Count  int
	Name   string
	Author string
}

//func NewBook(book *pb.Book) Book {
//	return Book{
//		ID:     int(book.GetID()),
//		ISBN:   book.GetISBN(),
//		Count:  int(book.GetCount()),
//		Name:   book.GetName(),
//		Author: book.GetAuthor(),
//	}
//}

func NewBooks(books []*pb.Book) []Book {
	var res []Book
	for _, book := range books {
		b := Book{
			ID:     int(book.GetID()),
			ISBN:   book.GetISBN(),
			Author: book.GetAuthor(),
			Name:   book.GetName(),
			Count:  int(book.GetCount()),
		}
		res = append(res, b)
	}
	return res
}

func NewFindBook(
	isbn string,
	name string,
	author string) *pb.FindBook {
	return &pb.FindBook{
		ISBN:   isbn,
		Name:   name,
		Author: author,
	}
}
