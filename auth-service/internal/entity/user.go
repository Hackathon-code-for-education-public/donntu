package entity

const (
	RoleApplicant = iota
	RoleStudent
	RoleManager
)

type Role int8

func (r Role) String() string {
	return [...]string{"applicant", "student", "manager"}[r]
}

type User struct {
	Id       string `json:"id" db:"id"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
	Role
}

func (u User) GetClaims() *UserClaims {
	return &UserClaims{
		Id:   u.Id,
		Role: u.Role,
	}
}

type UserClaims struct {
	Id   string `json:"id"`
	Role Role   `json:"role"`
}

type Tokens struct {
	Refresh string `json:"refresh_token"`
	Access  string `json:"access_token"`
}
