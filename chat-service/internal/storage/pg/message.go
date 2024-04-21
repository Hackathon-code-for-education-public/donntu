package pg

import (
	"chat-service/internal/models"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type MessageStorage struct {
	db *sqlx.DB
}

func NewMessageStorage(db *sqlx.DB) *MessageStorage {
	return &MessageStorage{db: db}
}

func (m MessageStorage) Save(ctx context.Context, msg *models.Message) error {
	sql, args, err := squirrel.
		Insert("messages").
		Columns("user_id", "chat_id", "text").
		Values(msg.UserId, msg.ChatId, msg.Text).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		slog.Error("failed build query", slog.String("err", err.Error()))
		return err
	}

	slog.Debug("exec query", slog.String("sql", sql))

	if _, err := m.db.ExecContext(ctx, sql, args...); err != nil {
		slog.Error("failed exec query", slog.String("err", err.Error()))
		return err
	}

	return nil
}

func (m MessageStorage) MarkAsRead(ctx context.Context, messageId int64) error {
	sql, args, err := squirrel.
		Update("messages").
		Set("read", true).
		Where(squirrel.Eq{"id": messageId}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		slog.Error("failed build query", slog.String("err", err.Error()))
		return err
	}

	slog.Debug("exec query", slog.String("sql", sql))
	if _, err := m.db.ExecContext(ctx, sql, args...); err != nil {
		slog.Error("failed exec query", slog.String("err", err.Error()))
		return err
	}
	return nil
}
