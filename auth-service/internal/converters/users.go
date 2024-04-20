package converters

import (
	"auth/internal/entity"
	"auth/internal/models"
)

func UserFromModelToEntity(model *models.User) *entity.User {

	var role entity.Role

	switch model.Role {
	case "applicant":
		role = entity.RoleApplicant
	case "student":
		role = entity.RoleStudent
	case "manager":
		role = entity.RoleManager
	}

	return &entity.User{
		Id:       model.Id,
		Email:    model.Email,
		Password: model.Password,
		Role:     role,
	}
}

func UserFromEntityToModel(entity *entity.User) *models.User {
	return &models.User{
		Id:       entity.Id,
		Email:    entity.Email,
		Password: entity.Password,
		Role:     entity.Role.String(),
	}
}
