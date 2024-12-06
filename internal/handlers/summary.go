package handlers

import (
	"fmt"

	"github.com/Isaac799/postgres-sql-migrations/internal/models"
	"github.com/Isaac799/postgres-sql-migrations/internal/services"
)

func SummarizeMigrations(cfg *models.Config) error {
	if err := services.SummarizeMigrations(cfg); err != nil {
		return fmt.Errorf("failed to summarize migrations: %w", err)
	}
	return nil
}
