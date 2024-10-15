package handlers

import (
	"fmt"
	"postgres_sql_migrations/internal/models"
	"postgres_sql_migrations/internal/services"
)

func ListAppliedMigrations(cfg *models.Config) error {
	if err := services.ListAppliedMigrations(cfg); err != nil {
		return fmt.Errorf("failed to list migrations: %w", err)
	}
	return nil
}
