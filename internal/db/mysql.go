package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/you/quizapi/internal/config"
)

func OpenMySQL(cfg config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", cfg.DSN)
	if err != nil {
		return nil, err
	}
	// warm-up ping
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
