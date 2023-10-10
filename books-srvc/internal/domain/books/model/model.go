package model

type Book struct {
	ID     int
	ISBN   string
	Count  int
	Name   string
	Author string
}

func NewFindBook(
	isbn string,
	name string,
	author string) Book {
	return Book{
		ISBN:   isbn,
		Name:   name,
		Author: author,
	}
}
