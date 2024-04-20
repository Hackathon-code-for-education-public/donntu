package grpc

import (
	"context"
	"verification-service/api/verification"
	"verification-service/internal/entity"
)

type Service interface {
	Send(ctx context.Context, req *entity.Verification) error
	Approve(ctx context.Context, requestId string) error
	Deny(ctx context.Context, requestId string, reason string) error
}

type Handler struct {
	verification.UnimplementedVerificationServer
	service Service
}

func New(service Service) *Handler {
	return &Handler{service: service}
}

func (h Handler) Send(ctx context.Context, request *verification.VerificationRequest) (*verification.SendRequestResponse, error) {
	v := &entity.Verification{
		UserId:          request.UserId,
		DocumentImageId: request.DocLink,
	}

	if err := h.service.Send(ctx, v); err != nil {
		return nil, err
	}

	return &verification.SendRequestResponse{
		RequestId: v.Id,
	}, nil
}

func (h Handler) Approve(ctx context.Context, request *verification.ApproveRequest) (*verification.Empty, error) {

	if err := h.service.Approve(ctx, request.RequestId); err != nil {
		return nil, err
	}

	return &verification.Empty{}, nil
}

func (h Handler) Decline(ctx context.Context, request *verification.DenialRequest) (*verification.Empty, error) {
	if err := h.service.Deny(ctx, request.RequestId, request.Reason); err != nil {
		return nil, err
	}

	return &verification.Empty{}, nil
}
