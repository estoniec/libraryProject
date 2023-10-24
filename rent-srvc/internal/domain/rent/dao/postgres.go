package dao

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"log/slog"
	"rent/internal/dal/postgres"
	"rent/internal/domain/rent/dto"
	psql "rent/pkg/postgresql"
)

type RentDAO struct {
	qb     sq.StatementBuilderType
	client psql.Client
}

func NewRentStorage(client psql.Client) *RentDAO {
	return &RentDAO{
		qb:     sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		client: client,
	}
}

func (repo *RentDAO) Create(ctx context.Context, dto dto.CreateDTO) error {
	sql, args, err := repo.qb.
		Insert(
			postgres.BooksUsersTable,
		).Columns(
		"fk_book_id",
		"fk_users_id",
		"returnAt",
	).Values(
		dto.BookID,
		dto.UserID,
		dto.ReturnAt,
	).ToSql()
	if err != nil {
		slog.Error(err.Error())
		return err
	}
	_, err = repo.client.Exec(ctx, sql, args...)
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	return nil
}
