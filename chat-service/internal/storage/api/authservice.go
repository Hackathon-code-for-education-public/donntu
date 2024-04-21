package api

import (
	"chat-service/api/auth"
	"chat-service/internal/config"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
)

type AuthService struct {
	client auth.AuthClient
}

func NewAuthService(cfg *config.Config) *AuthService {
	host := cfg.AuthService.Host
	port := cfg.AuthService.Port

	conn, err := grpc.NewClient(fmt.Sprintf("%s:%d", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	return &AuthService{
		client: auth.NewAuthClient(conn),
	}
}

func (s *AuthService) PatchRole(ctx context.Context, userId string) error {

	log := slog.With("userId", userId).With("service", "auth").With("method", "PatchRole")
	req := &auth.PatchRoleRequest{
		UserId:  userId,
		NewRole: auth.Role_student,
	}

	log.Debug("Request to update user role")
	if _, err := s.client.PatchRole(ctx, req); err != nil {
		log.Error("Failed to update user role: ", slog.String("error", err.Error()))
		return err
	}

	return nil
}
