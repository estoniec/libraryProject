package dto

type RentBookInput struct {
	BookID   int64
	UserID   int64
	ReturnAt int64
}

type RentBookOutput struct {
	Error  string
	Status int64
	ID     int64
}

type CreateDTO struct {
	BookID   int
	UserID   int
	ReturnAt int
}

func NewRentBookOutput(error string, status int64, id int64) RentBookOutput {
	return RentBookOutput{
		Error:  error,
		Status: status,
		ID:     id,
	}
}

func NewCreateDTO(bookID int, userID int, returnAt int) CreateDTO {
	return CreateDTO{
		BookID:   bookID,
		UserID:   userID,
		ReturnAt: returnAt,
	}
}
