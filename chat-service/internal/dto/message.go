package dto

type Message struct {
	Id     int64  `json:"id"`
	UserId string `json:"userId"`
	Text   string `json:"text"`
}
