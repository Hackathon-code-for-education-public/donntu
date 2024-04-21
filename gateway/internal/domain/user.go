package domain

import "gateway/api/auth"

type UserRole string

const (
	UserRoleManager    UserRole = "MANAGER"
	UserRoleUniversity UserRole = "UNIVERSITY"
	UserRoleApplicant  UserRole = "APPLICANT"
	UserRoleStudent    UserRole = "STUDENT"
)

type User struct {
	Email      string   `json:"email"`
	LastName   string   `json:"lastName"`
	FirstName  string   `json:"firstName"`
	MiddleName string   `json:"middleName"`
	Role       UserRole `json:"role"`
}

type UserClaims struct {
	Id   string   `json:"id"`
	Role UserRole `json:"role"`
}

func (u *User) ConvertRole() auth.Role {
	// TODO: Add more role
	switch u.Role {
	case UserRoleManager:
		return auth.Role_manager
	case UserRoleApplicant:
		return auth.Role_applicant
	case UserRoleStudent:
		return auth.Role_student
	}

	return auth.Role_applicant
}

func ConvertUserRole(role string) auth.Role {
	switch role {
	case "MANAGER":
		return auth.Role_manager
	case "APPLICANT":
		return auth.Role_applicant
	case "STUDENT":
		return auth.Role_student
	}

	return auth.Role_applicant
}

func ConvertRoleToGrpc(role UserRole) auth.Role {
	switch role {
	case "MANAGER":
		return auth.Role_manager
	case "APPLICANT":
		return auth.Role_applicant
	case "STUDENT":
		return auth.Role_student
	}

	return auth.Role_applicant
}

func ConvertUserRoleFromGrpc(role auth.Role) UserRole {
	switch role {
	case auth.Role_manager:
		return UserRoleManager
	case auth.Role_applicant:
		return UserRoleApplicant
	case auth.Role_student:
		return UserRoleStudent
	}

	return UserRoleApplicant
}
