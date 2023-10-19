package model

type Book struct {
	ID     int
	ISBN   string
	Count  int
	Name   string
	Author string
}

func NewBook(
	count int,
	isbn string,
	name string,
	author string) Book {
	return Book{
		Count:  count,
		ISBN:   isbn,
		Name:   name,
		Author: author,
	}
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
