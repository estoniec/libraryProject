package dao

import (
	"books-srvc/internal/dal/postgres"
	"books-srvc/internal/domain/books/dto"
	"books-srvc/internal/domain/books/model"
	psql "books-srvc/pkg/postgresql"
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v5"
	"log/slog"
)

type RegistrationDAO struct {
	qb     sq.StatementBuilderType
	client psql.Client
}

func NewBookStorage(client psql.Client) *RegistrationDAO {
	return &RegistrationDAO{
		qb:     sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		client: client,
	}
}

func (repo *RegistrationDAO) FindBy(ctx context.Context, dto dto.FindByInput) ([]model.Book, error) {
	where := sq.And{}
	if dto.Book.ISBN != "" {
		where = append(where, sq.Eq{"isbn": dto.Book.ISBN})
	}
	if dto.Book.Name != "" {
		where = append(where, sq.Eq{"name": dto.Book.Name})
	}
	if dto.Book.Author != "" {
		where = append(where, sq.Eq{"author": dto.Book.Author})
	}
	var books []model.Book
	sql, args, err := repo.qb.
		Select(
			"*",
		).From(
		postgres.BooksTable,
	).Where(
		where,
	).Limit(9).Offset(uint64(dto.Offset)).ToSql()
	if err != nil {
		slog.Error(err.Error())
		return []model.Book{}, err
	}
	rows, err := repo.client.Query(ctx, sql, args...)
	if err != nil {
		slog.Error(err.Error())
		if err == pgx.ErrNoRows {
			return []model.Book{}, fmt.Errorf("book is not found")
		}
		return []model.Book{}, err
	}

	for rows.Next() {
		var book model.Book
		err := rows.Scan(&book.ID, &book.ISBN, &book.Count, &book.Name, &book.Author)
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
