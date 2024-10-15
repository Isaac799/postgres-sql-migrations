package handlers

import (
	"fmt"
	"log"
	"postgres_sql_migrations/internal/models"
	"postgres_sql_migrations/internal/services"
)

func DropDatabase(cfg *models.Config, env string) error {
	exists, err := services.DatabaseExists(cfg)
	if err != nil {
		return fmt.Errorf("error checking database existence: %w", err)
	}
	if !exists {
		fmt.Printf("[ Notice ]: Database '%s' does not exist\n", cfg.Name)
		return nil
	}

	shouldDelete := confirmDropDatabase(cfg.Name, env)
	if !shouldDelete {
		return fmt.Errorf("database deletion aborted. Name did not match.")
	}

	if err := services.DeleteDatabase(cfg); err != nil {
		return fmt.Errorf("failed to delete database: %w", err)
	}
	return nil
}

func confirmDropDatabase(dbName string, env string) bool {
	fmt.Printf("WARNING: You are about to delete the database '%s' in '%s'. Type the database name to confirm: ", dbName, env)
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Fatalf("Failed to read input: %v", err)
	}
	return input == dbName
}
