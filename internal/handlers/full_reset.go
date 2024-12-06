package handlers

import (
	"fmt"

	"github.com/Isaac799/postgres-sql-migrations/internal/models"
)

func FullReset(cfg *models.Config, env string) error {
	if err := DropDatabase(cfg, env); err != nil {
		return fmt.Errorf("failed to drop database: %w", err)
	}

	if err := CreateDatabase(cfg); err != nil {
		return fmt.Errorf("failed to create database: %w", err)
	}

	if err := Migrate(cfg); err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}

	if err := SummarizeMigrations(cfg); err != nil {
		return fmt.Errorf("migration failed: %w", err)
	}

	return nil
}
