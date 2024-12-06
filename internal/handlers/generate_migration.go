package handlers

import (
	"fmt"

	"github.com/Isaac799/postgres-sql-migrations/internal/services"
)

func GenerateMigration(name string) error {
	if err := services.GenerateMigration(name); err != nil {
		return fmt.Errorf("failed to generate migration: %w", err)
	}
	return nil
}
