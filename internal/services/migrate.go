package services

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"sort"

	"github.com/Isaac799/postgres-sql-migrations/internal/models"
	"github.com/Isaac799/postgres-sql-migrations/internal/repository"
)

func DryRunMigrations(cfg *models.Config) error {
	db, err := repository.ConnectDB(cfg)
	if err != nil {
		return fmt.Errorf("error connecting to database: %w", err)
	}
	defer repository.CloseDB(db)

	migrations, err := getMigrationsToApply(db)
	if err != nil {
		return err
	}

	if len(migrations) == 0 {
		log.Println("Database is already up to date.")
		return nil
	}

	log.Println("Dry run migrations:")
	for _, migration := range migrations {
		log.Printf("Would apply migration: %s", migration)
	}

	return nil
}

func RunMigrations(cfg *models.Config) error {
	db, err := repository.ConnectDB(cfg)
	if err != nil {
		return fmt.Errorf("error connecting to database: %w", err)
	}
	defer repository.CloseDB(db)

	migrations, err := getMigrationsToApply(db)
	if err != nil {
		return err
	}

	if len(migrations) == 0 {
		log.Println("Database is already up to date.")
		return nil
	}

	return executeMigrations(db, migrations)
}

func getMigrationsToApply(db *sql.DB) ([]string, error) {
	files, err := os.ReadDir("./migrations")
	if err != nil {
		return nil, fmt.Errorf("failed to read migrations directory: %w", err)
	}

	var migrationsToApply []string

	// match 20060102150405_verb_noun
	re := regexp.MustCompile(`^\d{14}_[a-zA-Z0-9_]+\.sql$`)

	for _, file := range files {
		if !re.MatchString(file.Name()) {
			log.Printf("Skipped file (invalid format): %s", file.Name())
			continue
		}
		if migrationExists(db, file.Name()) {
			log.Printf("Skipped file (already applied): %s", file.Name())
			continue
		}
		migrationsToApply = append(migrationsToApply, file.Name())
	}

	// sort file names oldest -> newest
	sort.Strings(migrationsToApply)

	return migrationsToApply, nil
}

func executeMigrations(db *sql.DB, migrations []string) error {
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	for _, migration := range migrations {
		if err := applyMigration(tx, migration); err != nil {
			return err
		}
		log.Printf("Applied migration: %s", migration)
	}
	return nil
}

func applyMigration(tx *sql.Tx, migration string) error {
	path := filepath.Join("./migrations", migration)

	sqlFile, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open file %s: %w", migration, err)
	}
	defer sqlFile.Close()

	sqlBytes, err := io.ReadAll(sqlFile)
	if err != nil {
		return fmt.Errorf("failed to read file %s: %w", migration, err)
	}

	if _, err := tx.Exec(string(sqlBytes)); err != nil {
		return fmt.Errorf("failed to execute migration %s: %w", migration, err)
	}

	if err := recordMigration(tx, migration); err != nil {
		return fmt.Errorf("failed to record migration %s: %w", migration, err)
	}

	return nil
}

func migrationExists(db *sql.DB, migrationName string) bool {
	var count int
	query := `SELECT COUNT(*) FROM migrations WHERE name = $1`
	err := db.QueryRow(query, migrationName).Scan(&count)
	if err != nil {
		log.Printf("Error checking migration existence: %v", err)
		return false
	}
	return count > 0
}

func recordMigration(tx *sql.Tx, migrationName string) error {
	query := `INSERT INTO migrations (name, applied_at) VALUES ($1, NOW())`
	_, err := tx.Exec(query, migrationName)
	if err != nil {
		return fmt.Errorf("failed to record migration %s: %w", migrationName, err)
	}
	return nil
}
