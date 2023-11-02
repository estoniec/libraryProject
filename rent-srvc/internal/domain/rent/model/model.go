package model

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
