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

type Books_Users struct {
	ID       int64
	Users    User
	Books    Book
	isReturn bool
	isGet    bool
}
