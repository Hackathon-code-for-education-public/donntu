package engine

import "database/sql"

type DBEngine interface {
	GetDB() *sql.DB
	Close()
}
