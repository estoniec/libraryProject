package dao

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"log/slog"
	"project11/registration-svc/internal/dal/postgres"
	"project11/registration-svc/internal/domain/reg/dto"
	"project11/registration-svc/internal/domain/reg/model"
	psql "project11/registration-svc/pkg/postgresql"
)

type RegistrationDAO struct {
	qb     sq.StatementBuilderType
	client psql.Client
}

func NewRegistrationStorage(client psql.Client) *RegistrationDAO {
	return &RegistrationDAO{
		qb:     sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
		client: client,
	}
}

func (repo *RegistrationDAO) CreateUser(ctx context.Context, req model.User) error {
	sql, args, err := repo.qb.
		Insert(postgres.UserTable).
		Columns(
			"id",
			"phone",
			"username",
			"class",
			"status",
		).
		Values(
			req.ID,
			req.Phone,
			req.Username,
			req.Class,
			req.Status,
		).ToSql()
	if err != nil {
		slog.Error(err.Error())
		return err
	}

	_, execErr := repo.client.Exec(ctx, sql, args...)
	if execErr != nil {
		slog.Error(execErr.Error())

		return execErr
	}

	return nil
}

func (repo *RegistrationDAO) FindUser(ctx context.Context, req dto.CheckInput) (dto.CheckOutput, error) {
	var exists bool
	sql, args, err := repo.qb.
		Select(
			"*",
		).From(
		postgres.UserTable,
	).Where(
		sq.Eq{"id": req.ID},
	).Prefix(
		"SELECT EXISTS(",
	).Suffix(
		")",
	).ToSql()
	if err != nil {
		slog.Error(err.Error())
		return dto.NewCheckOutput(false), err
	}
	if err = repo.client.QueryRow(ctx, sql, args...).Scan(&exists); err != nil {
		slog.Error(err.Error())
		return dto.NewCheckOutput(false), err
	}

	return dto.NewCheckOutput(exists), nil
}
