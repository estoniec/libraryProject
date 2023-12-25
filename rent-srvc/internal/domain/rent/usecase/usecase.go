package usecase

import (
	"context"
	"rent/internal/domain/rent/dto"
)

type Service interface {
	ConfirmReturn(ctx context.Context, input dto.ConfirmReturnInput) (dto.ConfirmReturnOutput, error)
	EditCountBook(ctx context.Context, input dto.EditCountBookInput) (dto.EditCountBookOutput, error)
}

type Usecase struct {
	service Service
}

func NewUsecase(service Service) *Usecase {
	return &Usecase{
		service: service,
	}
}

func (u *Usecase) ConfirmReturn(ctx context.Context, input dto.ConfirmReturnInput) (dto.ConfirmReturnOutput, error) {

}
