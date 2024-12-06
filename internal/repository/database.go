package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Isaac799/postgres-sql-migrations/internal/models"

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

func DatabaseExists(cfg *models.Config) (bool, error) {
	db, err := sql.Open("postgres", cfg.ConnStr())
	if err != nil {
		return false, fmt.Errorf("failed to connect to PostgreSQL server: %w", err)
	}
	defer CloseDB(db)

	var exists bool
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = '%s');", cfg.Name)
	err = db.QueryRow(query).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error checking existence of database %s: %w", cfg.Name, err)
	}
	return exists, nil
}
