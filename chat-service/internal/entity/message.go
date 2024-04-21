package entity

import (
	"time"
)

type Message struct {
	Id        int64
	UserId    string
	TargetId  string
	Text      string
	Read      bool
	CreatedAt *time.Time
}
