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

func (repo *RegistrationDAO) FindByISBN(ctx context.Context, dto dto.FindByISBNInput) (model.Book, error) {
	var book model.Book
	sql, args, err := repo.qb.
		Select(
			"*",
		).From(
		postgres.BooksTable,
	).Where(
		sq.Eq{"isbn": dto.ISBN},
	).ToSql()
	if err != nil {
		slog.Error(err.Error())
		return model.Book{}, err
	}
	if err = repo.client.QueryRow(ctx, sql, args...).Scan(&book.ID, &book.ISBN, &book.Count, &book.Name, &book.Author); err != nil {
		slog.Error(err.Error())
		if err == pgx.ErrNoRows {
			return model.Book{}, fmt.Errorf("book is not found")
		}
		return model.Book{}, err
	}

	return book, nil
}

func (repo *RegistrationDAO) FindByAuthor(ctx context.Context, dto dto.FindByAuthorInput) ([]model.Book, error) {
	var books []model.Book
	sql, args, err := repo.qb.
		Select(
			"*",
		).From(
		postgres.BooksTable,
	).Where(
		sq.Eq{"author": dto.Author},
	).Limit(10).Offset(uint64(dto.Offset)).ToSql()
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

	return books, nil
}

func (repo *RegistrationDAO) FindByName(ctx context.Context, dto dto.FindByNameInput) ([]model.Book, error) {
	var books []model.Book
	sql, args, err := repo.qb.
		Select(
			"*",
		).From(
		postgres.BooksTable,
	).Where(
		sq.Eq{"name": dto.Name},
	).Limit(10).Offset(uint64(dto.Offset)).ToSql()
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

	return books, nil
}
