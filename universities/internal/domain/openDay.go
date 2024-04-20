package domain

import "time"

type OpenDay struct {
	UniversityName string
	Description    string
	Address        string
	Link           string
	Date           time.Time
}
