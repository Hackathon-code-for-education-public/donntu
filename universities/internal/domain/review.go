package domain

import "time"

type AuthorStatus string

var (
	StudentStatus  AuthorStatus = "Студент этого ВУЗ-а"
	GraduateStatus AuthorStatus = "Выпускник этого ВУЗ-а"
	ExpelledStatus AuthorStatus = "Отчисленный"
	SomeoneStatus  AuthorStatus = "Некто"
)

type Sentiment string

var (
	PositiveSentiment Sentiment = "Позитивный"
	NegativeSentiment Sentiment = "Негативный"
	NeutralSentiment  Sentiment = "Нейтральный"
)

type Review struct {
	Id           string       `json:"id"`
	UniversityId string       `json:"universityId"`
	AuthorStatus AuthorStatus `json:"authorStatus"`
	Sentiment    Sentiment    `json:"sentiment"`
	Date         time.Time    `json:"date"`
	Text         string       `json:"text"`
	RepliesCount int          `json:"repliesCount"`
	ParentId     *string      `json:"parentId"`
	AuthorId     string       `json:"authorId"`
}
