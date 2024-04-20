package services

import (
	"context"
	"github.com/google/uuid"
	"log/slog"
	"verification-service/internal/converters"
	"verification-service/internal/entity"
	"verification-service/internal/model"
)

type RequestStorage interface {
	Create(ctx context.Context, model *model.Verification) error
	Get(ctx context.Context, requestId string) (*model.Verification, error)
	PatchStatus(ctx context.Context, requestId string, status string) error
}

type ReasonStorage interface {
	Create(ctx context.Context, requestId string, reason string) error
}

type RoleUpdater interface {
	PatchRole(ctx context.Context, userId string) error
}

type Service struct {
	requestStorage RequestStorage
	reasonStorage  ReasonStorage
	roleUpdater    RoleUpdater
}

func New(requestStorage RequestStorage, reasonStorage ReasonStorage, roleUpdater RoleUpdater) *Service {
	return &Service{
		requestStorage: requestStorage,
		reasonStorage:  reasonStorage,
		roleUpdater:    roleUpdater,
	}
}

func (s Service) Send(ctx context.Context, v *entity.Verification) error {
	id, err := uuid.NewV7()
	if err != nil {
		return err
	}
	v.Id = id.String()

	m := converters.VerificationFromEntityToModel(v)

	if err := s.requestStorage.Create(ctx, m); err != nil {
		return err
	}

	return nil
}

func (s Service) Approve(ctx context.Context, requestId string) error {

	log := slog.With(slog.String("requestId", requestId), slog.String("method", "Approve"))

	log.Debug("patching status")
	if err := s.requestStorage.PatchStatus(ctx, requestId, entity.Status(entity.StatusApproved).String()); err != nil {
		return err
	}

	log.Debug("get request")
	v, err := s.requestStorage.Get(ctx, requestId)
	if err != nil {
		return err
	}

	log.Debug("patching role", slog.String("userId", v.UserId))
	if err := s.roleUpdater.PatchRole(ctx, v.UserId); err != nil {
		return err
	}

	return nil
}

func (s Service) Deny(ctx context.Context, requestId string, reason string) error {
	if err := s.requestStorage.PatchStatus(ctx, requestId, entity.Status(entity.StatusDenied).String()); err != nil {
		return err
	}

	if err := s.reasonStorage.Create(ctx, requestId, reason); err != nil {
		return err
	}

	return nil
}
