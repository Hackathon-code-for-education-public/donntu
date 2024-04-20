package grpc

import (
	"auth/api/auth"
	"auth/internal/config"
	"auth/internal/entity"
	"auth/internal/usecase"
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

var _ auth.AuthServer = (*Server)(nil)

type AuthUseCase interface {
	SignUp(ctx context.Context, user *entity.User) (*entity.Tokens, error)
	SignIn(ctx context.Context, user *entity.User) (*entity.Tokens, error)
	SingOut(ctx context.Context, accessToken string) error
	Authenticate(ctx context.Context, accessToken string, role entity.Role) (*entity.UserClaims, error)
	Refresh(ctx context.Context, refreshToken string) (*entity.Tokens, error)
}

type Server struct {
	cfg *config.Config

	uc AuthUseCase

	auth.UnimplementedAuthServer
}

func New(cfg *config.Config, uc AuthUseCase) *Server {
	return &Server{
		cfg: cfg,
		uc:  uc,
	}
}

func (s *Server) SignIn(ctx context.Context, request *auth.SignInRequest) (*auth.Tokens, error) {
	user := &entity.User{
		Email:    request.Email,
		Password: request.Password,
	}

	tokens, err := s.uc.SignIn(ctx, user)
	if err != nil {
		if errors.Is(err, usecase.ErrUserNotFound) {
			return nil, status.Error(codes.NotFound, "user not found")
		}

		if errors.Is(err, usecase.ErrInvalidCredentials) {
			return nil, status.Error(codes.InvalidArgument, "invalid credentials")
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &auth.Tokens{
		Access:  tokens.Access,
		Refresh: tokens.Refresh,
	}, nil
}

func (s *Server) SignUp(ctx context.Context, request *auth.SignUpRequest) (*auth.Tokens, error) {
	user := &entity.User{
		Email:    request.Email,
		Password: request.Password,
		Role:     entity.Role(request.Role),
	}

	tokens, err := s.uc.SignUp(ctx, user)
	if err != nil {
		if errors.Is(err, usecase.ErrUserAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, "user already exists")
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &auth.Tokens{
		Access:  tokens.Access,
		Refresh: tokens.Refresh,
	}, nil
}

func (s *Server) SignOut(ctx context.Context, request *auth.SignOutRequest) (*auth.Empty, error) {

	err := s.uc.SingOut(ctx, request.AccessToken)
	if err != nil {
		if errors.Is(err, usecase.ErrInvalidToken) {
			return nil, status.Error(codes.InvalidArgument, "invalid token")
		}

		if errors.Is(err, usecase.ErrSessionNotFound) {
			return nil, status.Error(codes.NotFound, "session not found")
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &auth.Empty{}, nil
}

func (s *Server) Auth(ctx context.Context, request *auth.AuthRequest) (*auth.AuthResponse, error) {

	claims, err := s.uc.Authenticate(ctx, request.AccessToken, entity.Role(request.Role))
	if err != nil {
		if errors.Is(err, usecase.ErrTokenExpired) {
			return nil, status.Error(codes.Unauthenticated, "token expired")
		}

		if errors.Is(err, usecase.ErrInvalidRole) {
			return nil, status.Error(codes.PermissionDenied, "invalid role")
		}

		if errors.Is(err, usecase.ErrSessionNotFound) {
			return nil, status.Error(codes.NotFound, "session not found")
		}

		if errors.Is(err, usecase.ErrInvalidToken) {
			return nil, status.Error(codes.InvalidArgument, "invalid token")
		}

		return nil, status.Error(codes.Internal, err.Error())
	}

	return &auth.AuthResponse{
		UserId: claims.Id,
		Role:   auth.Role(claims.Role),
	}, nil
}

func (s *Server) Refresh(ctx context.Context, request *auth.RefreshRequest) (*auth.Tokens, error) {

	log := ctx.Value("logger").(*slog.Logger).With("method", "Refresh")

	tokens, err := s.uc.Refresh(ctx, request.RefreshToken)
	if err != nil {
		if errors.Is(err, usecase.ErrInvalidToken) {
			log.Debug("invalid token")
			return nil, status.Error(codes.InvalidArgument, "invalid token")
		}

		if errors.Is(err, usecase.ErrSessionNotFound) {
			log.Debug("session not found")
			return nil, status.Error(codes.NotFound, "session not found")
		}

		log.Debug("internal server error", slog.String("err", err.Error()))
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &auth.Tokens{
		Access:  tokens.Access,
		Refresh: tokens.Refresh,
	}, nil
}
