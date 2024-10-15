package repository

import (
	"database/sql"
	"log"
	"postgres_sql_migrations/internal/models"

	_ "github.com/lib/pq"
)

func ConnectDB(cfg *models.Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.ConnStrDatabase())
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}

func CloseDB(db *sql.DB) {
	if err := db.Close(); err != nil {
		log.Printf("Error closing database connection: %v", err)
	}
}
