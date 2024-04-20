package model

import "time"

type Verification struct {
	Id              string     `db:"id"`
	UserId          string     `db:"user_id"`
	Status          string     `db:"status"`
	DocumentImageId string     `db:"doc_image_id"`
	CreatedAt       time.Time  `db:"created_at"`
	UpdatedAt       *time.Time `db:"updated_at"`
}
