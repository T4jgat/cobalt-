package postgres

import (
	"database/sql"
	"github.com/T4jgat/cobalt+/config"
	_ "github.com/lib/pq"
)

func New(cfg config.PG) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.URL)
	if err != nil {
		return nil, err
	}
	return db, nil
}
