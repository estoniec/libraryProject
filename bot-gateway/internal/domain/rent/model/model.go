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
	IsReturn bool
	IsGet    bool
	ReturnAt int64
}

func NewBooksUsers(models []*pb.BooksUsers) []BooksUsers {
	var res []BooksUsers
	for _, model := range models {
		b := BooksUsers{
			ID: model.GetID(),
			Books: Book{
				ID:     int(model.GetBook().GetID()),
				ISBN:   model.GetBook().GetISBN(),
				Name:   model.GetBook().GetName(),
				Author: model.GetBook().GetAuthor(),
				Count:  int(model.GetBook().GetCount()),
			},
			Users: User{
				ID:    model.GetUser().GetID(),
				Phone: model.GetUser().GetPhone(),
				Class: model.GetUser().GetClass(),
			},
			ReturnAt: model.GetReturnat(),
			IsReturn: model.GetIsReturn(),
			IsGet:    model.GetIsGet(),
		}
		res = append(res, b)
	}
	return res
}
