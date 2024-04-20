package pg

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"log/slog"
	"verification-service/internal/model"
)

type RequestStorage struct {
	db *sqlx.DB
}

func NewRequestStorage(db *sqlx.DB) *RequestStorage {
	return &RequestStorage{db: db}
}

func (r RequestStorage) Create(ctx context.Context, m *model.Verification) error {

	log := slog.With(slog.String("struct", "requestStorage"), slog.String("method", "Create"))

	sql, args, err := squirrel.
		Insert(requests_table).
		Columns("id", "user_id", "doc_image_id").
		Values(m.Id, m.UserId, m.DocumentImageId).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		log.Error("failed to build sql", slog.String("err", err.Error()))
		return nil
	}

	log.Debug("executing query", slog.String("sql", sql))
	if _, err := r.db.Exec(sql, args...); err != nil {
		log.Error("failed to execute query", slog.String("err", err.Error()))
		return nil
	}

	return nil
}

func (r RequestStorage) Get(ctx context.Context, id string) (*model.Verification, error) {
	log := slog.With(slog.String("struct", "requestStorage"), slog.String("method", "Get"), slog.String("id", id))

	sql, args, err := squirrel.Select("*").
		From(requests_table).
		Limit(1).
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		log.Error("failed build query", slog.String("err", err.Error()))
		return nil, err
	}

	log.Debug("query built", slog.String("sql", sql))
	v := new(model.Verification)

	if err := r.db.Get(v, sql, args...); err != nil {
		log.Error("failed query", slog.String("err", err.Error()))
		return nil, err
	}

	return v, nil
}

func (r RequestStorage) PatchStatus(ctx context.Context, requestId string, status string) error {
	log := slog.With(slog.String("struct", "requestStorage"), slog.String("method", "PatchStatus"))
	sql, args, err := squirrel.Update(requests_table).
		Set("status", status).
		Where(squirrel.Eq{"id": requestId}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		log.Error("failed to build query", slog.String("err", err.Error()))
		return err
	}

	log.Debug("query built", slog.String("sql", sql))

	if _, err := r.db.Exec(sql, args...); err != nil {
		return err
	}

	return nil
}
