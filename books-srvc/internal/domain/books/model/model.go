package model

type Book struct {
	ID     int
	ISBN   string
	Count  int
	Name   string
	Author string
}

//func NewBook(
//	id int,
//	isbn string,
//	count int,
//	name string,
//	author string) Book {
//	return Book{
//		ID:     id,
//		ISBN:   isbn,
//		Count:  count,
//		Name:   name,
//		Author: author,
//	}
//}
