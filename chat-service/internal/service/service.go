package service

import "context"

type ChatStorage interface {
	SendMessage(ctx context.Context)
}

type UserStorage interface {
}
