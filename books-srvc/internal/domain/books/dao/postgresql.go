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

type BooksDAO struct {
	qb     sq.StatementBuilderType
	client psql.Client
}

func NewBookStorage(client psql.Client) *BooksDAO {
	return &BooksDAO{
		qb:     sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		client: client,
	}
}

func (repo *BooksDAO) FindBy(ctx context.Context, dto dto.FindByInput) ([]model.Book, error) {
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
	if dto.Book.ID != 0 {
		where = append(where, sq.Eq{"id": dto.Book.ID})
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
		if book.Count > 0 {
			books = append(books, book)
		} else {
			continue
		}
	}
	if len(books) == 0 {
		return books, fmt.Errorf("book is not found")
	}
	return books, nil
}

func (repo *BooksDAO) Create(ctx context.Context, dto dto.CreateBookInput) (model.Book, error) {
	var book model.Book
	sql, args, err := repo.qb.
		Insert(
			postgres.BooksTable,
		).Columns(
		"isbn",
		"count",
		"name",
		"author",
	).Values(
		dto.Book.ISBN,
		dto.Book.Count,
		dto.Book.Name,
		dto.Book.Author,
	).Suffix("RETURNING *").ToSql()
	if err != nil {
		slog.Error(err.Error())
		return model.Book{}, err
	}
	err = repo.client.QueryRow(ctx, sql, args...).Scan(&book.ID, &book.ISBN, &book.Count, &book.Name, &book.Author)
	if err != nil {
		slog.Error(err.Error())
		if err == pgx.ErrNoRows {
			return model.Book{}, fmt.Errorf("book is not found")
		}
		return model.Book{}, err
	}

	return book, nil
}

func (repo *BooksDAO) EditCount(ctx context.Context, dto dto.EditCountBookInput) error {
	sql, args, err := repo.qb.
		Update(
			postgres.BooksTable,
		).Set(
		"count", dto.Count,
	).Where(
		sq.Eq{"isbn": dto.ISBN},
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

func (repo *BooksDAO) Delete(ctx context.Context, dto dto.DeleteBookInput) error {
	sql, args, err := repo.qb.
		Delete(
			postgres.BooksTable,
		).Where(
		sq.Eq{"isbn": dto.ISBN},
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
