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

func (repo *RentDAO) FindBy(ctx context.Context, dto dto.FindByInput) ([]model.BooksUsers, error) {
	where := sq.And{}
	if dto.ID != 0 {
		where = append(where, sq.Eq{"id": dto.ID})
		where = append(where, sq.Eq{"isreturn": false})
	}
	if dto.Time != 0 {
		where = append(where, sq.Eq{"isreturn": false})
		where = append(where, sq.Eq{"isget": true})
		where = append(where, sq.LtOrEq{"returnat": dto.Time})
	}
	if dto.UserID != 0 && dto.BookID != 0 {
		where = append(where, sq.Eq{"fk_users_id": dto.UserID})
		where = append(where, sq.Eq{"fk_book_id": dto.BookID})
		where = append(where, sq.Eq{"isreturn": false})
	}
	var books []model.BooksUsers
	sql, args, err := repo.qb.
		Select(
			"id",
			"returnat",
			"isreturn",
			"isget",
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
		where,
	).Join(
		"public.books ON public.books_users.fk_book_id = public.books.book_id",
	).Join(
		"public.users ON public.books_users.fk_users_id = public.users.user_id",
	).Limit(5).Offset(uint64(dto.Offset)).ToSql()
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
		var book model.BooksUsers
		err := rows.Scan(&book.ID, &book.ReturnAt, &book.IsReturn, &book.IsGet, &book.Books.ID, &book.Books.ISBN, &book.Books.Count, &book.Books.Name, &book.Books.Author, &book.Users.ID, &book.Users.Phone, &book.Users.Class)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	if len(books) == 0 {
		return books, fmt.Errorf("book is not found")
	}
	return books, nil
}
