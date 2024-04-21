package models

type Chat struct {
	Id      int64 `db:"id"`
	User1Id int64
	User2Id int64
}
