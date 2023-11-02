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
		Book: &pb.BooksUsers{
			Book: &pb.Book{
				ID:     int64(output.Book.ID),
				Name:   output.Book.Books.Name,
				Author: output.Book.Books.Author,
				ISBN:   output.Book.Books.ISBN,
				Count:  int64(output.Book.Books.Count),
			},
			User: &pb.User{
				ID:    output.Book.Users.ID,
				Phone: output.Book.Users.Phone,
			},
		},
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

func NewGetDebtInput(req *pb.GetDebtRequest) dto.GetDebtInput {
	return dto.GetDebtInput{
		Time: req.GetTime(),
	}
}

func NewGetDebtOutput(output dto.GetDebtOutput) *pb.GetDebtResponse {
	var debts []*pb.BooksUsers
	if len(output.Debt) == 0 {
		debts = nil
	} else {
		debts = make([]*pb.BooksUsers, len(output.Debt))
		for i, debt := range output.Debt {
			debts[i] = &pb.BooksUsers{
				ID: debt.ID,
				Book: &pb.Book{
					ID:     int64(debt.Books.ID),
					ISBN:   debt.Books.ISBN,
					Name:   debt.Books.Name,
					Author: debt.Books.Author,
					Count:  int64(debt.Books.Count),
				},
				User: &pb.User{
					ID:    debt.Users.ID,
					Phone: debt.Users.Phone,
					Class: debt.Users.Class,
				},
			}
		}
	}
	return &pb.GetDebtResponse{
		Error:  output.Error,
		Status: output.Status,
		Debt:   debts,
	}
}

func NewFindByUidAndBidInput(req *pb.FindByUidAndBidRequest) dto.FindByUIDAndBIDInput {
	return dto.FindByUIDAndBIDInput{
		Uid: req.GetUID(),
		Bid: req.GetBID(),
	}
}

func NewFindByUidAndBidOutput(output dto.FindByUIDAndBIDOutput) *pb.FindByUidAndBidResponse {
	return &pb.FindByUidAndBidResponse{
		Error:  output.Error,
		Status: output.Status,
		Id:     output.ID,
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

func FindByInput(req *pb.FindByRequest) dto.FindByInput {
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
		Book:   models,
	}
}
