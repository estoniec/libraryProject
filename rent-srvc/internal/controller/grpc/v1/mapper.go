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

func NewConfirmRentInput(req *pb.ConfirmRentRequest) dto.ConfirmRentInput {
	return dto.ConfirmRentInput{
		ID: req.GetID(),
	}
}

func NewConfirmRentOutput(output dto.ConfirmRentOutput) *pb.ConfirmRentResponse {
	return &pb.ConfirmRentResponse{
		Error:  output.Error,
		Status: output.Status,
	}
}

func NewConfirmReturnInput(req *pb.ConfirmReturnRequest) dto.ConfirmReturnInput {
	return dto.ConfirmReturnInput{
		ID: req.GetID(),
	}
}

func NewConfirmReturnOutput(output dto.ConfirmReturnOutput) *pb.ConfirmReturnResponse {
	return &pb.ConfirmReturnResponse{
		Error:  output.Error,
		Status: output.Status,
	}
}

func NewFindByInput(req *pb.FindByRequest) dto.FindByInput {
	return dto.FindByInput{
		Offset: req.GetOffset(),
		ID:     req.GetID(),
		Time:   req.GetTime(),
		UserID: req.GetUserID(),
		BookID: req.GetBookID(),
	}
}

func NewFindByOutput(output dto.FindByOutput) *pb.FindByResponse {
	var models []*pb.BooksUsers
	if len(output.Model) == 0 {
		models = nil
	} else {
		models = make([]*pb.BooksUsers, len(output.Model))
		for i, model := range output.Model {
			models[i] = &pb.BooksUsers{
				ID:       model.ID,
				IsReturn: model.IsReturn,
				IsGet:    model.IsGet,
				Returnat: model.ReturnAt,
				Book: &pb.Book{
					ID:     int64(model.Books.ID),
					ISBN:   model.Books.ISBN,
					Count:  int64(model.Books.Count),
					Name:   model.Books.Name,
					Author: model.Books.Author,
				},
				User: &pb.User{
					ID:    model.Users.ID,
					Phone: model.Users.Phone,
					Class: model.Users.Class,
				},
			}
		}
	}
	return &pb.FindByResponse{
		Status: output.Status,
		Error:  output.Error,
		Model:  models,
	}
}
