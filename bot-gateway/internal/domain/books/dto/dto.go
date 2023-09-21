package dto

type FindByISBNInput struct {
	ISBN string
}

func NewISBNInput(isbn string) FindByISBNInput {
	return FindByISBNInput{
		ISBN: isbn,
	}
}
