package handlers

import (
	"fmt"
	"postgres_sql_migrations/internal/models"
	"postgres_sql_migrations/internal/services"
)

func SummarizeMigrations(cfg *models.Config) error {
	if err := services.SummarizeMigrations(cfg); err != nil {
		return fmt.Errorf("failed to summarize migrations: %w", err)
	}
	return nil
}
