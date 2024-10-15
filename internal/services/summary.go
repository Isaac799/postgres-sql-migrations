package services

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"postgres_sql_migrations/internal/models"
	"postgres_sql_migrations/internal/repository"
	"strings"
)

func SummarizeMigrations(cfg *models.Config) error {
	db, err := repository.ConnectDB(cfg)
	if err != nil {
		return fmt.Errorf("error connecting to database: %w", err)
	}
	defer repository.CloseDB(db)

	localMigrations, err := getLocalMigrations()
	if err != nil {
		return fmt.Errorf("failed to get local migrations: %w", err)
	}

	appliedMigrations, err := getAppliedMigrations(db)
	if err != nil {
		return fmt.Errorf("failed to get applied migrations: %w", err)
	}

	migrationsToApply, err := getMigrationsToApply(db)
	if err != nil {
		return fmt.Errorf("failed to get migrations to apply: %w", err)
	}

	// tracks local migrations by their names
	localMigrationNames := make(map[string]int)
	for _, name := range localMigrations {
		localMigrationNames[name] = 0
	}

	// identify migrations that are not part of the local migrations
	var nonLocalCount int
	var nonLocalMigrations []string
	for _, name := range appliedMigrations {
		if _, exists := localMigrationNames[name]; !exists {
			nonLocalCount++
			nonLocalMigrations = append(nonLocalMigrations, name)
		}
	}

	var summaryBuilder strings.Builder
	summaryBuilder.WriteString("\nMigration Summary:\n")

	appliedCount := len(appliedMigrations)
	summaryBuilder.WriteString(fmt.Sprintf("\nApplied Migrations: %d\n", appliedCount))
	for _, name := range appliedMigrations {
		summaryBuilder.WriteString(fmt.Sprintf(" - %s\n", name))
	}

	localCount := len(localMigrations)
	summaryBuilder.WriteString(fmt.Sprintf("\nLocal Migrations: %d\n", localCount))
	for _, name := range localMigrations {
		summaryBuilder.WriteString(fmt.Sprintf(" - %s\n", name))
	}

	toApplyCount := len(migrationsToApply)
	summaryBuilder.WriteString(fmt.Sprintf("\nUnapplied Local Migrations: %d\n", toApplyCount))
	for _, name := range migrationsToApply {
		summaryBuilder.WriteString(fmt.Sprintf(" - %s\n", name))
	}

	if nonLocalCount > 0 {
		summaryBuilder.WriteString(fmt.Sprintf("\n[ Warning ]: Non-Local Applied Migrations: %d\n", nonLocalCount))
		for _, name := range nonLocalMigrations {
			summaryBuilder.WriteString(fmt.Sprintf(" - %s\n", name))
		}
	}

	if toApplyCount > 0 {
		summaryBuilder.WriteString("\n[ Notice ]: You have migrations to apply. Consider running '-migrate'\n")
	}

	log.Println(summaryBuilder.String())

	return nil
}

func getLocalMigrations() ([]string, error) {
	files, err := os.ReadDir("./migrations")
	if err != nil {
		return nil, fmt.Errorf("failed to read migrations directory: %w", err)
	}

	var migrationsToApply []string
	for _, file := range files {
		if filepath.Ext(file.Name()) != ".sql" {
			continue
		}
		migrationsToApply = append(migrationsToApply, file.Name())
	}

	return migrationsToApply, nil
}

func getAppliedMigrations(db *sql.DB) ([]string, error) {
	rows, err := db.Query("SELECT name FROM migrations ORDER BY applied_at ASC")
	if err != nil {
		return nil, fmt.Errorf("failed to query migrations: %w", err)
	}
	defer rows.Close()

	var appliedMigrations []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		appliedMigrations = append(appliedMigrations, name)
	}

	return appliedMigrations, nil
}
