package services

import (
	"database/sql"
	"fmt"
	"log"
	"postgres_sql_migrations/internal/models"
	"postgres_sql_migrations/internal/repository"
)

func CreateDatabase(cfg *models.Config) error {
	if err := createDatabase(cfg); err != nil {
		return err
	}

	if err := setupMigrations(cfg); err != nil {
		return err
	}

	return nil
}

func createDatabase(cfg *models.Config) error {
	db, err := sql.Open("postgres", cfg.ConnStr())
	if err != nil {
		return fmt.Errorf("failed to connect to PostgreSQL server: %w", err)
	}
	defer repository.CloseDB(db)

	_, err = db.Exec(fmt.Sprintf("CREATE DATABASE %s;", cfg.Name))
	if err != nil {
		return fmt.Errorf("failed to create database %s: %w", cfg.Name, err)
	}
	log.Printf("Database %s created successfully.", cfg.Name)

	return nil
}

func setupMigrations(cfg *models.Config) error {
	db, err := sql.Open("postgres", cfg.ConnStrDatabase())
	if err != nil {
		return fmt.Errorf("failed to connect to new database %s: %w", cfg.Name, err)
	}
	defer repository.CloseDB(db)

	if err := setupMigrationsTable(db); err != nil {
		return fmt.Errorf("failed to set up migrations table: %w", err)
	}

	return nil
}

func setupMigrationsTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS migrations (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		applied_at TIMESTAMP NOT NULL DEFAULT NOW()
	);
	`
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to create migrations table: %w", err)
	}
	log.Println("Migrations table created or already exists.")
	return nil
}
