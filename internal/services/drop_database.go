package services

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Isaac799/postgres-sql-migrations/internal/models"
	"github.com/Isaac799/postgres-sql-migrations/internal/repository"
)

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
