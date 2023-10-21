package dto

type RentBookInput struct {
	BookID   int64
	UserID   int64
	ReturnAt int64
}

type RentBookOutput struct {
	Author string
	Name   string
	Error  string
	Status int64
}

type CreateDTO struct {
	BookID   int
	UserID   int
	ReturnAt int
}

func NewRentBookOutput(author string, name string, error string, status int64) RentBookOutput {
	return RentBookOutput{
		Author: author,
		Name:   name,
		Error:  error,
		Status: status,
	}
}

func NewCreateDTO(bookID int, userID int, returnAt int) CreateDTO {
	return CreateDTO{
		BookID:   bookID,
		UserID:   userID,
		ReturnAt: returnAt,
	}
}
