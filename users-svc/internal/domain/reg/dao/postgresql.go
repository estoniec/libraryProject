package dao

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"log/slog"
	"registration-svc/internal/dal/postgres"
	"registration-svc/internal/domain/reg/dto"
	"registration-svc/internal/domain/reg/model"
	psql "registration-svc/pkg/postgresql"
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
			"user_id",
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

func (repo *RegistrationDAO) FindUserByRole(ctx context.Context, req dto.CheckRoleInput) (dto.CheckRoleOutput, error) {
	var role int
	sql, args, err := repo.qb.
		Select(
			"status",
		).From(
		postgres.UserTable,
	).Where(
		sq.Eq{"user_id": req.ID},
	).ToSql()
	if err != nil {
		slog.Error(err.Error())
		return dto.NewCheckRoleOutput(0, err.Error(), 404), err
	}
	if err = repo.client.QueryRow(ctx, sql, args...).Scan(&role); err != nil {
		slog.Error(err.Error())
		return dto.NewCheckRoleOutput(0, err.Error(), 404), err
	}

	return dto.NewCheckRoleOutput(role, "", 200), nil
}
