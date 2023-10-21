package dao

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
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

func (repo *RentDAO) Create(ctx context.Context, dto dto.CreateDTO) (int, error) {
	var res int
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
	).Suffix("RETURNING \"id\"").ToSql()
	if err != nil {
		slog.Error(err.Error())
		return res, err
	}
	err = repo.client.QueryRow(ctx, sql, args...).Scan(&res)
	if err != nil {
		slog.Error(err.Error())
		if err == pgx.ErrNoRows {
			return res, fmt.Errorf("book is not found")
		}
		return res, err
	}

	return res, nil
}
