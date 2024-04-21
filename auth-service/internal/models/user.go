package models

import "time"

type User struct {
	Id         string     `db:"id"`
	Email      string     `db:"email"`
	Password   string     `db:"password"`
	Role       string     `db:"role"`
	CreatedAt  time.Time  `db:"created_at"`
	UpdatedAt  *time.Time `db:"updated_at"`
	LastName   string     `db:"last_name"`
	FirstName  string     `db:"first_name"`
	MiddleName string     `db:"middle_name"`
}
