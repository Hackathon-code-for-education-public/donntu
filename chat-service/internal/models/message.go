package models

import "time"

type Message struct {
	Id        int64      `db:"id"`
	UserId    string     `db:"user_id"`
	ChatId    string     `db:"chat_id"`
	Text      string     `db:"text"`
	Read      bool       `db:"read"`
	CreatedAt *time.Time `db:"created_at"`
}
