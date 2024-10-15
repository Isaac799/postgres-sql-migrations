package services

import (
	"database/sql"
	"fmt"
	"log"
	"postgres_sql_migrations/internal/models"
	"postgres_sql_migrations/internal/repository"
)

func DatabaseExists(cfg *models.Config) (bool, error) {
	db, err := sql.Open("postgres", cfg.ConnStr())
	if err != nil {
		return false, fmt.Errorf("failed to connect to PostgreSQL server: %w", err)
	}
	defer repository.CloseDB(db)

	var exists bool
	query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM pg_database WHERE datname = '%s');", cfg.Name)
	err = db.QueryRow(query).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error checking existence of database %s: %w", cfg.Name, err)
	}
	return exists, nil
}

func DeleteDatabase(cfg *models.Config) error {
	db, err := sql.Open("postgres", cfg.ConnStr())
	if err != nil {
		return fmt.Errorf("failed to connect to PostgreSQL server: %w", err)
	}
	defer repository.CloseDB(db)

	_, err = db.Exec(fmt.Sprintf("DROP DATABASE %s;", cfg.Name))
	if err != nil {
		return fmt.Errorf("failed to delete database %s: %w", cfg.Name, err)
	}
	log.Printf("Database %s deleted successfully.", cfg.Name)
	return nil
}
