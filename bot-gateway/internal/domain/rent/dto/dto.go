package rentService

type RentBookOutput struct {
	Error  string
	Status int64
}

type RentBookDTO struct {
	BookID   int64
	UserID   int64
	ReturnAt int64
}

func NewRentBookOutput(error string, status int64) RentBookOutput {
	return RentBookOutput{
		Error:  error,
		Status: status,
	}
}

func NewRentBookDTO(bookID int64, userID int64, returnAt int64) RentBookDTO {
	return RentBookDTO{
		BookID:   bookID,
		UserID:   userID,
		ReturnAt: returnAt,
	}
}
