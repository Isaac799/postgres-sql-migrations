package handlers

import (
	"fmt"
	"postgres_sql_migrations/internal/models"
	"postgres_sql_migrations/internal/services"
)

func CreateDatabase(cfg *models.Config) error {
	exists, err := services.DatabaseExists(cfg)
	if err != nil {
		return fmt.Errorf("error checking database existence: %w", err)
	}
	if exists {
		fmt.Printf("[ Notice ]: Database '%s' already exists\n", cfg.Name)
		return nil
	}

	if err := services.CreateDatabase(cfg); err != nil {
		return fmt.Errorf("failed to create database: %w", err)
	}
	return nil
}
