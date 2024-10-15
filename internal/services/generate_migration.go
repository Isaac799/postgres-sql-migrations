package services

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func GenerateMigration(name string) error {
	filename := createMigrationFilename(name)
	migrationsDir := "./migrations"

	if err := ensureMigrationsDir(migrationsDir); err != nil {
		return err
	}

	filePath := filepath.Join(migrationsDir, filename)
	if err := createMigrationFile(filePath, name); err != nil {
		return err
	}

	log.Printf("Migration file created: %s", filePath)
	return nil
}

func createMigrationFilename(name string) string {
	timestamp := time.Now().Format("20060102150405")
	return fmt.Sprintf("%s_%s.sql", timestamp, name)
}

func ensureMigrationsDir(dir string) error {
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create migrations directory: %w", err)
	}
	return nil
}

func createMigrationFile(filePath, name string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create migration file %s: %w", filePath, err)
	}
	defer file.Close()

	migrationTemplate := "-- Migration: %s\n-- Created on: %s\n\n-- TODO: Write migration SQL\n"
	if _, err := file.WriteString(fmt.Sprintf(migrationTemplate, name, time.Now().Format(time.RFC1123))); err != nil {
		return fmt.Errorf("failed to write to migration file %s: %w", filePath, err)
	}

	return nil
}
