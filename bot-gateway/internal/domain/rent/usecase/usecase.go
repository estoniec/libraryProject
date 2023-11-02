package rentService

import (
	"context"
	dto2 "gateway/internal/controller/telegram/dto"
	dto "gateway/internal/domain/rent/dto"
	rentService "gateway/internal/domain/rent/model"
	pb "github.com/estoniec/libraryProject/contracts/gen/go/books_users"
	"strconv"
)

type Usecase struct {
	client pb.BooksUsersServiceClient
}

func NewUsecase(client pb.BooksUsersServiceClient) *Usecase {
	return &Usecase{
		client: client,
	}
}

func (u *Usecase) RentBook(ctx context.Context, input dto2.RentInput) (dto.RentBookOutput, error) {
	bookIDint, err := strconv.Atoi(input.BookID)
	if err != nil {
		return dto.NewRentBookOutput(err.Error(), 404, 0), err
	}
	dtoRent := dto.NewRentBookDTO(int64(bookIDint), input.UserID, input.ReturnAt)
	res, err := u.client.RentBook(ctx, &pb.RentBookRequest{
		BookID:   dtoRent.BookID,
		UserID:   dtoRent.UserID,
		ReturnAt: dtoRent.ReturnAt,
	})
	if err != nil {
		return dto.NewRentBookOutput(err.Error(), 404, 0), err
	}
	return dto.NewRentBookOutput(res.GetError(), res.GetStatus(), res.GetId()), nil
}

func (u *Usecase) FindBook(ctx context.Context, input dto2.FindBookInput) (dto.FindBookOutput, error) {
	res, err := u.client.FindBy(ctx, &pb.FindByRequest{
		Id: input.ID,
	})
	if err != nil {
		return dto.NewFindBookOutput(err.Error(), 404, nil), err
	}
	return dto.NewFindBookOutput(res.GetError(), res.GetStatus(), res.GetBook()), nil
	res, err := u.client.FindBook(ctx, &pb.FindBookRequest{
		Id: input.ID,
	})
	if err != nil {
		return dto.NewFindBookOutput(err.Error(), 404, nil), err
	}
	return dto.NewFindBookOutput(res.GetError(), res.GetStatus(), res.GetBook()), nil
}

func (u *Usecase) ConfirmRent(ctx context.Context, input dto2.ConfirmRentInput) (dto.ConfirmRentOutput, error) {
	res, err := u.client.ConfirmRent(ctx, &pb.ConfirmRentRequest{
		ID: input.ID,
	})
	if err != nil {
		return dto.NewConfirmRentOutput(err.Error(), 404), err
	}
	return dto.NewConfirmRentOutput(res.GetError(), res.GetStatus()), nil
}

func (u *Usecase) GetDebt(ctx context.Context, input dto2.GetDebtInput) (dto.GetDebtOutput, error) {
	res, err := u.client.GetDebt(ctx, &pb.GetDebtRequest{
		Time: input.Time,
	})
	debts := rentService.NewBooksUsers(res.GetDebt())
	if err != nil {
		return dto.NewGetDebtOutput(err.Error(), 404, debts), err
	}
	return dto.NewGetDebtOutput(res.GetError(), res.GetStatus(), debts), nil
}

func (u *Usecase) CheckRent(ctx context.Context, input dto2.CheckRentInput) (dto.CheckRentOutput, error) {
	res, err := u.client.FindByUidAndBid(ctx, &pb.FindByUidAndBidRequest{
		UID: input.UID,
		BID: input.BID,
	})
	if err != nil && err.Error() != "rpc error: code = Unknown desc = rent is not found" {
		return dto.NewCheckRentOutput(err.Error(), 404, 0), err
	}
	return dto.NewCheckRentOutput(res.GetError(), res.GetStatus(), res.GetId()), nil
}

func (u *Usecase) ConfirmReturn(ctx context.Context, input dto2.ConfirmReturnInput) (dto.ConfirmReturnOutput, error) {
	res, err := u.client.ConfirmReturn(ctx, &pb.ConfirmReturnRequest{
		ID: input.ID,
	})
	if err != nil {
		return dto.NewConfirmReturnOutput(err.Error(), 404), err
	}
	return dto.NewConfirmReturnOutput(res.GetError(), res.GetStatus()), nil
}
