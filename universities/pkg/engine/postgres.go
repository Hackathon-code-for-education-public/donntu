package engine

import (
	"database/sql"
	"fmt"
	"time"
)

type postgres struct {
	connAttempts int
	connTimeout  time.Duration

	db *sql.DB
}

type DBConnString string

func NewPostgresDB(url DBConnString, connAttempts int, connTimeout time.Duration) (DBEngine, error) {
	pg := &postgres{
		connAttempts: connAttempts,
		connTimeout:  connTimeout,
	}

	var err error
	for pg.connAttempts > 0 {
		pg.db, err = sql.Open("postgres", string(url))
		if err != nil {
			return nil, fmt.Errorf("erro while connecting to db: %v", err)
		}

		time.Sleep(pg.connTimeout)
		pg.connAttempts--
	}

	return pg, nil
}

func (p *postgres) GetDB() *sql.DB {
	return p.db
}

func (p *postgres) Close() {
	if p.db != nil {
		p.db.Close()
	}
}
