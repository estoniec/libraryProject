package dto

type RentBookInput struct {
	BookID   int64
	UserID   int64
	ReturnAt int64
}

type RentBookOutput struct {
	ID     int
	Error  string
	Status int64
}

type CreateDTO struct {
	BookID   int
	UserID   int
	ReturnAt int
}

func NewRentBookOutput(id int, error string, status int64) RentBookOutput {
	return RentBookOutput{
		ID:     id,
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
