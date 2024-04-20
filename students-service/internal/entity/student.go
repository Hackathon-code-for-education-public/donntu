package entity

type Student struct {
	Id           string `json:"id"`
	UniversityId string `json:"university"`
}

type Review struct {
	Id           string `json:"id"`
	StudentId    string `json:"student"`
	UniversityId string `json:"university"`
	Body         string `json:"body"`
	Rating       int    `json:"rating"`
}
