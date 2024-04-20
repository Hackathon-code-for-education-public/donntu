package pg

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type ReasonStorage struct {
	db *sqlx.DB
}

func NewReasonStorage(db *sqlx.DB) *ReasonStorage {
	return &ReasonStorage{db: db}
}

func (r ReasonStorage) Create(ctx context.Context, requestId string, reason string) error {
	sql, args, err := squirrel.
		Insert(reasons_table).
		Columns("request_id", "reason").
		Values(requestId, reason).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil
	}

	if _, err := r.db.Exec(sql, args...); err != nil {
		return nil
	}

	return nil
}
