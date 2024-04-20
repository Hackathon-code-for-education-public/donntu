package pg

import (
	"auth/internal/models"
	"auth/internal/usecase"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

var _ usecase.UserStorage = (*UserStorage)(nil)

type UserStorage struct {
	db *sqlx.DB
}

func NewUserStorage(db *sqlx.DB) *UserStorage {
	return &UserStorage{db: db}
}

func (u *UserStorage) Get(ctx context.Context, id string) (*models.User, error) {

	log := ctx.Value("logger").(*slog.Logger).With("method", "Get")

	var user models.User

	query, args, err := squirrel.Select("*").
		From(CREDENTIALS_TABLE).
		Where(squirrel.Eq{"id": id}).
		Limit(1).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		log.Error("Failed to generate SQL query", slog.String("err", err.Error()))
		return nil, err
	}

	log.Debug("executing query", slog.String("query", query), slog.Any("args", args))

	err = u.db.Get(&user, query, args...)
	if err != nil {
		log.Error("Failed to execute query", slog.String("err", err.Error()))
		return nil, err
	}

	return &user, nil
}

func (u *UserStorage) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	log := ctx.Value("logger").(*slog.Logger).With("method", "GetByEmail")
	var user models.User
	query, args, err := squirrel.Select("*").
		From(CREDENTIALS_TABLE).
		Where(squirrel.Eq{"email": email}).
		Limit(1).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		log.Error("Failed to generate SQL query", slog.String("err", err.Error()))
		return nil, err
	}
	log.Debug("executing query", slog.String("query", query), slog.Any("args", args))
	if err := u.db.Get(&user, query, args...); err != nil {
		log.Error("Failed to execute query", slog.String("err", err.Error()))
		return nil, err

	}

	return &user, nil
}

func (u *UserStorage) Create(ctx context.Context, user *models.User) error {

	log := ctx.Value("logger").(*slog.Logger).With("method", "Create")

	query, args, err := squirrel.Insert(CREDENTIALS_TABLE).
		Columns("id", "email", "password", "role").
		Values(user.Id, user.Email, user.Password, user.Role).
		Suffix("RETURNING *").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		log.Error("failed to generate SQL query", slog.String("err", err.Error()))
		return err
	}

	log.Debug("executing query", slog.String("query", query), slog.Any("args", args))

	if err := u.db.Get(user, query, args...); err != nil {
		log.Error("failed to execute query", slog.String("err", err.Error()))
		return err
	}

	log.Debug("query result user", slog.Any("user", user))
	return nil
}

func (u *UserStorage) Update(ctx context.Context, user *models.User) error {
	log := ctx.Value("logger").(*slog.Logger).With("method", "Update")

	query, args, err := squirrel.Update("users").
		Set("email", user.Email).
		Set("password", user.Password).
		Where(squirrel.Eq{"id": user.Id}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		log.Error("failed to generate SQL query", slog.String("err", err.Error()))
		return err
	}

	log.Debug("executing query", slog.String("query", query), slog.Any("args", args))
	if err := u.db.Get(user, query, args...); err != nil {
		log.Error("failed to execute query", slog.String("err", err.Error()))
		return err
	}

	log.Debug("query result user", slog.Any("user", user))
	return nil
}

func (u *UserStorage) Delete(ctx context.Context, id int) error {
	log := ctx.Value("logger").(*slog.Logger).With("method", "Delete")
	query, args, err := squirrel.Delete("users").
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()

	if err != nil {
		log.Error("failed to generate SQL query", slog.String("err", err.Error()))
		return err
	}

	log.Debug("executing query", slog.String("query", query), slog.Any("args", args))
	if _, err := u.db.Exec(query, args...); err != nil {
		log.Error("failed to execute query", slog.String("err", err.Error()))
		return err
	}

	log.Debug("successfully deleted user", slog.Any("id", id))

	return nil
}

func (u *UserStorage) PatchRole(ctx context.Context, id string, role string) error {
	log := ctx.Value("logger").(*slog.Logger).With("method", "PatchRole")

	query, args, err := squirrel.Update("users").
		Set("role", role).
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		log.Error("failed to generate SQL query", slog.String("err", err.Error()))
		return err
	}

	log.Debug("executing query", slog.String("query", query), slog.Any("args", args))
	if _, err := u.db.Exec(query, args...); err != nil {
		log.Error("failed to execute query", slog.String("err", err.Error()))
		return err
	}

	log.Debug("successfully updated user", slog.Any("id", id))
	return nil
}
