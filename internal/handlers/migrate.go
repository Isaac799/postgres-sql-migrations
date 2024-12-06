package handlers

import (
	"fmt"

	"github.com/Isaac799/postgres-sql-migrations/internal/models"
	"github.com/Isaac799/postgres-sql-migrations/internal/services"
)

func Migrate(cfg *models.Config) error {
	if err := services.RunMigrations(cfg); err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}
	return nil
}

func MigrateDryRun(cfg *models.Config) error {
	if err := services.DryRunMigrations(cfg); err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}
	return nil
}
