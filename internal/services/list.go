package services

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/Isaac799/postgres-sql-migrations/internal/models"
	"github.com/Isaac799/postgres-sql-migrations/internal/repository"
)

func ListAppliedMigrations(cfg *models.Config) error {
	db, err := repository.ConnectDB(cfg)
	if err != nil {
		return fmt.Errorf("error connecting to database: %w", err)
	}
	defer repository.CloseDB(db)

	rows, err := db.Query("SELECT name, applied_at FROM migrations ORDER BY applied_at ASC")
	if err != nil {
		return fmt.Errorf("failed to query migrations: %w", err)
	}
	defer rows.Close()

	var builder strings.Builder
	builder.WriteString("\nApplied Migrations:\n")
	for rows.Next() {
		var name string
		var appliedAt time.Time
		if err := rows.Scan(&name, &appliedAt); err != nil {
			return fmt.Errorf("failed to scan row: %w", err)
		}
		builder.WriteString(fmt.Sprintf(" - %s (applied at: %s)\n", name, appliedAt.Format(time.RFC1123)))
	}

	log.Println(builder.String())
	return nil
}
