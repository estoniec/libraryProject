package service

import (
	"context"
	"rent/internal/domain/rent/dto"
	"rent/internal/domain/rent/model"
)

type Repository interface {
	Create(ctx context.Context, dto dto.CreateDTO) (int, error)
	Find(ctx context.Context, dto dto.FindBookInput) (model.BooksUsers, error)
	UpdateGet(ctx context.Context, dto dto.ConfirmRentInput) error
	FindByTime(ctx context.Context, dto dto.GetDebtInput) ([]model.BooksUsers, error)
	FindByUIDAndBID(ctx context.Context, dto dto.FindByUIDAndBIDInput) (int64, error)
	UpdateReturn(ctx context.Context, dto dto.ConfirmReturnInput) error
	FindAll(ctx context.Context, dto dto.FindBookInput) (model.BooksUsers, error)
	FindBy(ctx context.Context, dto dto.FindByInput) ([]model.BooksUsers, error)
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) RentBook(ctx context.Context, input dto.RentBookInput) (dto.RentBookOutput, error) {
	dtoRent := dto.NewCreateDTO(int(input.BookID), int(input.UserID), int(input.ReturnAt))
	id, err := s.repository.Create(ctx, dtoRent)
	if err != nil {
		return dto.NewRentBookOutput(err.Error(), 404, 0), err
	}
	return dto.NewRentBookOutput("", 200, int64(id)), nil
}

func (s *Service) FindBook(ctx context.Context, input dto.FindBookInput) (dto.FindBookOutput, error) {
	book, err := s.repository.Find(ctx, input)
	if err != nil {
		return dto.NewFindBookOutput(err.Error(), 404, book), err
	}
	return dto.NewFindBookOutput("", 200, book), nil
}

func (s *Service) ConfirmRent(ctx context.Context, input dto.ConfirmRentInput) (dto.ConfirmRentOutput, error) {
	err := s.repository.UpdateGet(ctx, input)
	if err != nil {
		return dto.NewConfirmRentOutput(err.Error(), 404), err
	}
	return dto.NewConfirmRentOutput("", 200), nil
}

func (s *Service) GetDebt(ctx context.Context, input dto.GetDebtInput) (dto.GetDebtOutput, error) {
	debt, err := s.repository.FindByTime(ctx, input)
	if err != nil {
		return dto.NewGetDebtOutput(err.Error(), 404, []model.BooksUsers{}), err
	}
	return dto.NewGetDebtOutput("", 200, debt), nil
}

func (s *Service) FindByUidAndBid(ctx context.Context, input dto.FindByUIDAndBIDInput) (dto.FindByUIDAndBIDOutput, error) {
	id, err := s.repository.FindByUIDAndBID(ctx, input)
	if err != nil {
		return dto.NewFindByUIDAndBIDOutput(err.Error(), 404, 0), err
	}
	return dto.NewFindByUIDAndBIDOutput("", 200, id), nil
}

func (s *Service) ConfirmReturn(ctx context.Context, input dto.ConfirmReturnInput) (dto.ConfirmReturnOutput, error) {
	err := s.repository.UpdateReturn(ctx, input)
	if err != nil {
		return dto.NewConfirmReturnOutput(err.Error(), 404), err
	}
	return dto.NewConfirmReturnOutput("", 200), nil
}

func (s *Service) FindBy(ctx context.Context, input dto.FindByInput) (dto.FindByOutput, error) {
	model, err := s.repository.FindBy(ctx, input)
	if err != nil {
		return dto.NewFindByOutput(err.Error(), 404, model), err
	}
	return dto.NewFindByOutput("", 200, model), nil
}
