package handlers

import (
	"fmt"
	"postgres_sql_migrations/internal/models"
	"postgres_sql_migrations/internal/services"
)

func CreateDatabase(cfg *models.Config) error {
	if err := services.CreateDatabase(cfg); err != nil {
		return fmt.Errorf("failed to create database: %w", err)
	}
	return nil
}
