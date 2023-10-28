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

func NewFindBookInput(req *pb.FindBookRequest) dto.FindBookInput {
	return dto.FindBookInput{
		ID: req.GetId(),
	}
}

func NewFindBookOutput(output dto.FindBookOutput) *pb.FindBookResponse {
	return &pb.FindBookResponse{
		Error:  output.Error,
		Status: output.Status,
		Book: &pb.Book{
			ID:     int64(output.Book.ID),
			Name:   output.Book.Name,
			Author: output.Book.Author,
			ISBN:   output.Book.ISBN,
			Count:  int64(output.Book.Count),
		},
	}
}
