package domain

import (
	"context"
	"gateway/pkg/file"
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
		CreatePanorama(ctx context.Context, panorama *Panorama) (*Panorama, error)
		GetPanoramas(ctx context.Context, universityId string, category string) ([]*Panorama, error)

		GetUniversitiesTop(ctx context.Context, n int) ([]*University, error)
		GetUniversities(ctx context.Context, offset, limit int) ([]*University, error)
		GetUniversity(ctx context.Context, universityID string) (*University, error)
		SearchUniversities(ctx context.Context, name string) ([]*University, error)
	}

	FileService interface {
		Upload(ctx context.Context, reader file.Reader) (string, error)
	}
)
