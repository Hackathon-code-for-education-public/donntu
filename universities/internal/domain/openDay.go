package domain

import "time"

type OpenDay struct {
	UniversityName string    `json:"universityName"`
	Description    string    `json:"description"`
	Address        string    `json:"address"`
	Link           string    `json:"link"`
	Date           time.Time `json:"date"`
}
