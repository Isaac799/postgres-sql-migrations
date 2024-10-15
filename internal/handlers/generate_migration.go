package handlers

import (
	"fmt"
	"postgres_sql_migrations/internal/services"
)

func GenerateMigration(name string) error {
	if err := services.GenerateMigration(name); err != nil {
		return fmt.Errorf("failed to generate migration: %w", err)
	}
	return nil
}
