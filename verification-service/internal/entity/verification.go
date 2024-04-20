package entity

import "time"

type Verification struct {
	Id              string     `json:"id"`
	UserId          string     `json:"user_id"`
	Status          Status     `json:"status"`
	DocumentImageId string     `json:"documentImageId"`
	CreatedAt       time.Time  `json:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at"`
}

type Reason struct {
	RequestId string `json:"userId"`
	Reason    string `json:"reason"`
}

type Status int32

const (
	StatusPending = iota
	StatusApproved
	StatusDenied
	StatusUnknown
)

func NewStatus(s string) Status {
	switch s {
	case "pending":
		return StatusPending
	case "approved":
		return StatusApproved
	case "denied":
		return StatusDenied
	}
	return StatusUnknown
}

func (v Status) String() string {
	return [...]string{"pending", "approved", "denied", "unknown"}[v]
}
