package converters

import (
	"verification-service/internal/entity"
	"verification-service/internal/model"
)

func VerificationFromEntityToModel(e *entity.Verification) *model.Verification {
	return &model.Verification{
		Id:              e.Id,
		UserId:          e.UserId,
		Status:          e.Status.String(),
		DocumentImageId: e.DocumentImageId,
		CreatedAt:       e.CreatedAt,
		UpdatedAt:       e.UpdatedAt,
	}
}

func VerificationFromModelToEntity(m *model.Verification) *entity.Verification {
	return &entity.Verification{
		Id:              m.Id,
		UserId:          m.UserId,
		Status:          entity.NewStatus(m.Status),
		DocumentImageId: m.DocumentImageId,
		CreatedAt:       m.CreatedAt,
		UpdatedAt:       m.UpdatedAt,
	}
}
