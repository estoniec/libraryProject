package v1

import (
	pb "github.com/estoniec/libraryProject/contracts/gen/go/books_users"
	"rent/internal/domain/rent/dto"
)

func NewRentBookInput(req *pb.RentBookRequest) dto.RentBookInput {
	return dto.RentBookInput{
		BookID:   req.GetBookID(),
		UserID:   req.GetUserID(),
		ReturnAt: req.GetReturnAt(),
	}
}

func NewRentBookOutput(output dto.RentBookOutput) *pb.RentBookResponse {
	return &pb.RentBookResponse{
		Error:  output.Error,
		Status: output.Status,
		Id:     output.ID,
	}
}
