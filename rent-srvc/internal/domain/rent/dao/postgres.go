package dao

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"log"
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

func (repo *RentDAO) Find(ctx context.Context, dto dto.FindBookInput) (model.Book, error) {
	var book model.Book
	sql, args, err := repo.qb.
		Select(
			"public.books.book_id",
			"public.books.isbn",
			"public.books.count",
			"public.books.name",
			"public.books.author",
		).From(
		postgres.BooksUsersTable,
	).Where(
		sq.Eq{"id": dto.ID}).Join(
		"public.books ON public.books_users.fk_book_id = public.books.book_id",
	).ToSql()
	if err != nil {
		slog.Error(err.Error())
		return model.Book{}, err
	}
	err = repo.client.QueryRow(ctx, sql, args...).Scan(&book.ID, &book.ISBN, &book.Count, &book.Name, &book.Author)
	if err != nil {
		slog.Error(err.Error())
		return model.Book{}, err
	}
	log.Print(book.ID, book.ISBN, book.Count, book.Author, book.Name)
	return book, nil
}
