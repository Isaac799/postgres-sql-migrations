package handlers

import (
	"fmt"

	"github.com/Isaac799/postgres-sql-migrations/internal/models"
	"github.com/Isaac799/postgres-sql-migrations/internal/repository"
	"github.com/Isaac799/postgres-sql-migrations/internal/services"
)

func CreateDatabase(cfg *models.Config) error {
	exists, err := repository.DatabaseExists(cfg)
	if err != nil {
		return err
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
