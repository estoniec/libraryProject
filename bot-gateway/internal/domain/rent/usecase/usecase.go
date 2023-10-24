package rentService

import (
	"context"
	dto2 "gateway/internal/controller/telegram/dto"
	dto "gateway/internal/domain/rent/dto"
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
		return dto.NewRentBookOutput(err.Error(), 404), err
	}
	dtoRent := dto.NewRentBookDTO(int64(bookIDint), input.UserID, input.ReturnAt)
	res, err := u.client.RentBook(ctx, &pb.RentBookRequest{
		BookID:   dtoRent.BookID,
		UserID:   dtoRent.UserID,
		ReturnAt: dtoRent.ReturnAt,
	})
	if err != nil {
		return dto.NewRentBookOutput(err.Error(), 404), err
	}
	return dto.NewRentBookOutput(res.GetError(), res.GetStatus()), nil
}
