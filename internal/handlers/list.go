package handlers

import (
	"fmt"

	"github.com/Isaac799/postgres-sql-migrations/internal/models"
	"github.com/Isaac799/postgres-sql-migrations/internal/services"
)

func ListAppliedMigrations(cfg *models.Config) error {
	if err := services.ListAppliedMigrations(cfg); err != nil {
		return fmt.Errorf("failed to list migrations: %w", err)
	}
	return nil
}
