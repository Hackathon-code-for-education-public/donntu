package usecase

import (
	"auth/internal/converters"
	"auth/internal/entity"
	"context"
)

func (u *UseCase) GetUser(ctx context.Context, id string) (*entity.User, error) {

	m, err := u.userStorage.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return converters.UserFromModelToEntity(m), nil
}

func (u *UseCase) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {
	m, err := u.userStorage.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	return converters.UserFromModelToEntity(m), nil
}

func (u *UseCase) DeleteUser(ctx context.Context, id string) error {
	return u.userStorage.Delete(ctx, id)
}
