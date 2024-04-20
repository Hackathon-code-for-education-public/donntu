package domain

import "time"

type Review struct {
	UniversityId string    `json:"universityId"`
	AuthorStatus string    `json:"authorStatus"`
	Sentiment    string    `json:"sentiment"`
	Date         time.Time `json:"date"`
	Text         string    `json:"text"`
	RepliesCount int       `json:"repliesCount"`
}
