package domain

import (
	"context"
)

type (
	AuthService interface {
		SignIn(ctx context.Context, credentials *Credentials) (*Tokens, error)
		SignUp(ctx context.Context, user *User, password string) (*Tokens, error)
		SignOut(ctx context.Context, accessToken string) error
		Verify(ctx context.Context, accessToken string, role *UserRole) (*UserClaims, error)
		Refresh(ctx context.Context, refreshToken string) (*Tokens, error)
		GetUser(ctx context.Context, userId string) (*User, error)
	}

	UniversityService interface {
		GetOpenDays(ctx context.Context, universityId string) ([]*OpenDay, error)
		GetReviews(ctx context.Context, universityId string, offset int, limit int) ([]*Review, error)
	}
)
