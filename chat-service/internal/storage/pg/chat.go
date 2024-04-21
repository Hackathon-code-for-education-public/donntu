package pg

import (
	"chat-service/internal/dto"
	"chat-service/internal/models"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"log/slog"
)

type ChatStorage struct {
	db *sqlx.DB
}

func NewChatStorage(db *sqlx.DB) *ChatStorage {
	return &ChatStorage{
		db: db,
	}
}

func (c ChatStorage) Create(ctx context.Context, id string, userId string, targetId string) error {

	tx, err := c.db.Beginx()
	if err != nil {
		return err
	}

	sql, args, err := squirrel.
		Insert("chats").
		Columns("chat_id").
		Values(id).
		ToSql()
	if err != nil {
		slog.Error("failed build insert into chats error", slog.String("err", err.Error()))
		return err
	}

	slog.Debug("executing sql", slog.String("sql", sql))
	if _, err := tx.Exec(sql, args...); err != nil {
		slog.Error("failed insert into chats error", slog.String("err", err.Error()))
		return err
	}

	sql, args, err = squirrel.
		Insert("users_to_chats").
		Columns("chat_id", "user_id").
		Values(id, userId).
		ToSql()
	if err != nil {
		slog.Error("failed build insert into users_to_chats error user_id", slog.String("err", err.Error()))
		return err
	}

	slog.Debug("executing sql", slog.String("sql", sql))
	if _, err := tx.Exec(sql, args...); err != nil {
		if err := tx.Rollback(); err != nil {
			slog.Error("rollback error", slog.String("err", err.Error()))
			return err
		}
		slog.Error("insert error", slog.String("err", err.Error()))
		return err
	}

	sql, args, err = squirrel.
		Insert("users_to_chats").
		Columns("chat_id", "user_id").
		Values(id, targetId).
		ToSql()
	if err != nil {
		slog.Error("failed build insert into users_to_chats error target_id", slog.String("err", err.Error()))
		if err := tx.Rollback(); err != nil {
			slog.Error("rollback error", slog.String("err", err.Error()))
			return err
		}
		return err
	}

	slog.Debug("executing sql", slog.String("sql", sql))

	if _, err := tx.Exec(sql, args...); err != nil {
		if err := tx.Rollback(); err != nil {
			slog.Error("rollback error", slog.String("err", err.Error()))
			return err
		}
		slog.Error("insert error in ", slog.String("err", err.Error()))
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (c ChatStorage) ListChatByUserId(ctx context.Context, userId string) ([]*dto.Chat, error) {

	type row struct {
		ChatId string `db:"chat_id"`
		UserId string `db:"user_id"`
	}

	sql, args, err := squirrel.
		Select("*").
		From("users_to_chat uc").
		Where(squirrel.Eq{"uc.user_id": userId}).
		ToSql()
	if err != nil {
		slog.Error("failed build sql", slog.String("err", err.Error()))
		return nil, err
	}

	slog.Debug("executing sql", slog.String("sql", sql))

	rows := make([]*row, 0)

	if err := c.db.Select(rows, sql, args...); err != nil {
		slog.Error("failed executing sql", slog.String("sql", sql))
		return nil, err
	}

	chch := make([]*dto.Chat, 0, len(rows))
	for _, r := range rows {
		chch = append(chch, &dto.Chat{
			Id:     r.ChatId,
			UserId: r.UserId,
		})
	}

	return chch, err
}

func (c ChatStorage) GetHistory(ctx context.Context, chatId string) ([]*models.Message, error) {
	sql, args, err := squirrel.
		Select("*").
		From("messages").
		Where(squirrel.Eq{"chat_id": chatId}).
		ToSql()
	if err != nil {
		slog.Error("failed build sql", slog.String("err", err.Error()))
		return nil, err
	}
	slog.Debug("executing sql", slog.String("sql", sql))
	rows := make([]*models.Message, 0)
	if err := c.db.Select(rows, sql, args...); err != nil {
		slog.Error("failed executing sql", slog.String("sql", sql))
		return nil, err
	}

	return rows, nil
}
