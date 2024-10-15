package handlers

import (
	"fmt"
	"postgres_sql_migrations/internal/models"
	"postgres_sql_migrations/internal/services"
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
