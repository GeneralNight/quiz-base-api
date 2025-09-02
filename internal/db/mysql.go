package db

import (
	"errors"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"quiz-base-api/internal/config"
)

var dbInstance *sqlx.DB

func GetDB() (*sqlx.DB, error) {
	if dbInstance == nil {
		return nil, errors.New("DB not initialized")
	}
	return dbInstance, nil
}

func OpenMySQL(cfg config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("mysql", cfg.DSN)
	if err != nil {
		return nil, err
	}
	// warm-up ping
	if err := db.Ping(); err != nil {
		return nil, err
	}
	dbInstance = db
	return db, nil
}
