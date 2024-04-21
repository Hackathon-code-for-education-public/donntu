package domain

import "time"

type Review struct {
	ReviewId     string    `json:"reviewId"`
	UniversityId string    `json:"universityId"`
	AuthorStatus string    `json:"authorStatus"`
	Sentiment    string    `json:"sentiment"`
	Date         time.Time `json:"date"`
	Text         string    `json:"text"`
	RepliesCount int       `json:"repliesCount"`
	ParentId     *string   `json:"parentId,omitempty"`
	AuthorId     *string   `json:"authorId,omitempty"`
}
