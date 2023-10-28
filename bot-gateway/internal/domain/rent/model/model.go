package rentService

import pb "github.com/estoniec/libraryProject/contracts/gen/go/books_users"

type User struct {
	ID       int64
	Phone    string
	Username string
	Class    string
	Status   int
}

type Book struct {
	ID     int
	ISBN   string
	Count  int
	Name   string
	Author string
}

type BooksUsers struct {
	ID       int64
	Users    User
	Books    Book
	isReturn bool
	isGet    bool
}

func NewBooksUsers(debts []*pb.BooksUsers) []BooksUsers {
	var res []BooksUsers
	for _, debt := range debts {
		b := BooksUsers{
			ID: debt.GetID(),
			Books: Book{
				ID:     int(debt.GetBook().GetID()),
				ISBN:   debt.GetBook().GetISBN(),
				Name:   debt.GetBook().GetName(),
				Author: debt.GetBook().GetAuthor(),
				Count:  int(debt.GetBook().GetCount()),
			},
			Users: User{
				ID:    debt.GetUser().GetID(),
				Phone: debt.GetUser().GetPhone(),
				Class: debt.GetUser().GetClass(),
			},
		}
		res = append(res, b)
	}
	return res
}
