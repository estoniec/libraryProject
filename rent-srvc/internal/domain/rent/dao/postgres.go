package dao

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"log/slog"
	"rent/internal/dal/postgres"
	"rent/internal/domain/rent/dto"
	"rent/internal/domain/rent/model"
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
	var id int
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
		return id, err
	}
	err = repo.client.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		slog.Error(err.Error())
		return id, err
	}
	return id, nil
}

func (repo *RentDAO) Find(ctx context.Context, dto dto.FindBookInput) (model.BooksUsers, error) {
	var book model.BooksUsers
	sql, args, err := repo.qb.
		Select(
			"public.books.book_id",
			"public.books.isbn",
			"public.books.count",
			"public.books.name",
			"public.books.author",
			"public.users.user_id",
			"public.users.phone",
		).From(
		postgres.BooksUsersTable,
	).Where(
		sq.Eq{"id": dto.ID},
		sq.Eq{"isreturn": false},
	).Join(
		"public.books ON public.books_users.fk_book_id = public.books.book_id",
	).Join(
		"public.users ON public.books_users.fk_users_id = public.users.user_id",
	).ToSql()
	if err != nil {
		slog.Error(err.Error())
		return model.BooksUsers{}, err
	}
	err = repo.client.QueryRow(ctx, sql, args...).Scan(&book.Books.ID, &book.Books.ISBN, &book.Books.Count, &book.Books.Name, &book.Books.Author, &book.Users.ID, &book.Users.Phone)
	if err != nil {
		slog.Error(err.Error())
		return model.BooksUsers{}, err
	}
	return book, nil
}

func (repo *RentDAO) UpdateGet(ctx context.Context, dto dto.ConfirmRentInput) error {
	sql, args, err := repo.qb.
		Update(
			postgres.BooksUsersTable,
		).Set(
		"isget",
		true,
	).Where(
		sq.Eq{"id": dto.ID}).ToSql()
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

func (repo *RentDAO) FindByTime(ctx context.Context, dto dto.GetDebtInput) ([]model.BooksUsers, error) {
	var debts []model.BooksUsers
	sql, args, err := repo.qb.
		Select(
			"id",
			"public.books.book_id",
			"public.books.isbn",
			"public.books.count",
			"public.books.name",
			"public.books.author",
			"public.users.user_id",
			"public.users.phone",
			"public.users.class",
		).From(
		postgres.BooksUsersTable,
	).Where(
		sq.And{
			sq.Eq{"isreturn": false},
			sq.Eq{"isget": true},
			sq.LtOrEq{"returnat": dto.Time},
		}).Join(
		"public.books ON public.books_users.fk_book_id = public.books.book_id",
	).Join(
		"public.users ON public.books_users.fk_users_id = public.users.user_id",
	).ToSql()
	if err != nil {
		slog.Error(err.Error())
		return []model.BooksUsers{}, err
	}

	rows, err := repo.client.Query(ctx, sql, args...)
	if err != nil {
		slog.Error(err.Error())
		if err == pgx.ErrNoRows {
			return []model.BooksUsers{}, fmt.Errorf("book is not found")
		}
		return []model.BooksUsers{}, err
	}

	for rows.Next() {
		var debt model.BooksUsers
		err := rows.Scan(&debt.ID, &debt.Books.ID, &debt.Books.ISBN, &debt.Books.Count, &debt.Books.Name, &debt.Books.Author, &debt.Users.ID, &debt.Users.Phone, &debt.Users.Class)
		if err != nil {
			return nil, err
		}
		debts = append(debts, debt)
	}
	if len(debts) == 0 {
		return debts, fmt.Errorf("book is not found")
	}
	return debts, nil
}

func (repo *RentDAO) FindByUIDAndBID(ctx context.Context, dto dto.FindByUIDAndBIDInput) (int64, error) {
	var id int64
	sql, args, err := repo.qb.
		Select(
			"id",
		).From(
		postgres.BooksUsersTable,
	).Where(
		sq.And{
			sq.Eq{"fk_users_id": dto.Uid},
			sq.Eq{"fk_book_id": dto.Bid},
			sq.Eq{"isreturn": false},
		}).ToSql()
	if err != nil {
		slog.Error(err.Error())
		return 0, err
	}

	err = repo.client.QueryRow(ctx, sql, args...).Scan(&id)
	if err != nil {
		slog.Error(err.Error())
		if err == pgx.ErrNoRows {
			return 0, fmt.Errorf("rent is not found")
		}
		return 0, err
	}

	return id, nil
}

func (repo *RentDAO) UpdateReturn(ctx context.Context, dto dto.ConfirmReturnInput) error {
	sql, args, err := repo.qb.
		Update(
			postgres.BooksUsersTable,
		).Set(
		"isreturn",
		true,
	).Where(
		sq.Eq{"id": dto.ID}).ToSql()
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
