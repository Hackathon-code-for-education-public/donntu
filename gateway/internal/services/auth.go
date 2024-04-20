package services

import (
	"context"
	"errors"
	"fmt"
	"gateway/api/auth"
	"gateway/internal/config"
	"gateway/internal/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

var (
	ErrUnauthorized   = errors.New("unauthorized")
	ErrForbidden      = errors.New("forbidden")
	ErrNotFound       = errors.New("not found")
	ErrInvalidRequest = errors.New("invalid request")
)

type AuthService struct {
	cfg    *config.Config
	client auth.AuthClient
}

func NewAuthService(cfg *config.Config) domain.AuthService {
	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", cfg.Services.AuthService.Host, cfg.Services.AuthService.Port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := auth.NewAuthClient(conn)

	return &AuthService{
		cfg:    cfg,
		client: client,
	}
}

func (s *AuthService) SignIn(ctx context.Context, credentials *domain.Credentials) (*domain.Tokens, error) {
	req := &auth.SignInRequest{
		Email:    credentials.Email,
		Password: credentials.Password,
	}

	res, err := s.client.SignIn(ctx, req)
	if err != nil {
		return nil, err
	}

	return &domain.Tokens{
		Access:  res.Access,
		Refresh: res.Refresh,
	}, nil
}

func (s *AuthService) SignUp(ctx context.Context, user *domain.User, password string) (*domain.Tokens, error) {
	req := &auth.SignUpRequest{
		Email:    user.Email,
		Password: password,
		Role:     user.ConvertRole(),
	}
	fmt.Println(req)

	res, err := s.client.SignUp(ctx, req)
	if err != nil {
		return nil, err
	}

	return &domain.Tokens{
		Access:  res.Access,
		Refresh: res.Refresh,
	}, nil
}

func (s *AuthService) SignOut(ctx context.Context, accessToken string) error {
	req := &auth.SignOutRequest{
		AccessToken: accessToken,
	}

	if _, err := s.client.SignOut(ctx, req); err != nil {
		return err
	}

	return nil
}

func (s *AuthService) Refresh(ctx context.Context, refreshToken string) (*domain.Tokens, error) {
	req := &auth.RefreshRequest{
		RefreshToken: refreshToken,
	}

	res, err := s.client.Refresh(ctx, req)
	if err != nil {
		return nil, err
	}

	return &domain.Tokens{
		Access:  res.Access,
		Refresh: res.Refresh,
	}, nil
}

func (s *AuthService) Verify(ctx context.Context, accessToken string, role string) (*domain.UserClaims, error) {
	req := &auth.AuthRequest{
		AccessToken: accessToken,
		Role:        domain.ConvertUserRole(role),
	}

	res, err := s.client.Auth(ctx, req)
	if err != nil {
		if e, ok := status.FromError(err); ok {
			switch e.Code() {
			case codes.Unauthenticated:
				return nil, ErrUnauthorized
			case codes.PermissionDenied:
				return nil, ErrForbidden
			case codes.NotFound:
				return nil, ErrNotFound
			case codes.InvalidArgument:
				return nil, ErrInvalidRequest
			default:
				return nil, err
			}
		}

		return nil, err
	}

	return &domain.UserClaims{
		Id:   res.UserId,
		Role: domain.ConvertUserRoleFromGrpc(res.Role),
	}, nil
}
