package router

import (
	"flag"

	"github.com/Isaac799/postgres-sql-migrations/internal/models"
)

func parseCliFlags() models.Flags {
	var flags models.Flags

	helpTooltip := "Show help information"
	envTooltip := "Environment (from your config.json e.g. dev)"
	genMigrationTooltip := "Generate a new migration file with the given name"
	listMigrationsTooltip := "List available migrations"
	createTooltip := "Creates the database (initial setup)"
	dropTooltip := "Drops the database (removes all data)"
	migrateTooltip := "Run migrations"
	migrateDryTooltip := "Dry run migrations without applying changes"
	fullResetTooltip := "Drops, creates, and migrates the database"
	summaryTooltip := "Summarize migrations"

	flag.BoolVar(&flags.Help, "help", false, helpTooltip)
	flag.BoolVar(&flags.Help, "h", false, helpTooltip)
	flag.StringVar(&flags.Env, "env", "", envTooltip)
	flag.StringVar(&flags.Env, "e", "", envTooltip)
	flag.StringVar(&flags.GenMigration, "generate-migration", "", genMigrationTooltip)
	flag.StringVar(&flags.GenMigration, "gm", "", genMigrationTooltip)
	flag.BoolVar(&flags.ListMigrations, "list-migrations", false, listMigrationsTooltip)
	flag.BoolVar(&flags.ListMigrations, "lm", false, listMigrationsTooltip)
	flag.BoolVar(&flags.Create, "create", false, createTooltip)
	flag.BoolVar(&flags.Create, "c", false, createTooltip)
	flag.BoolVar(&flags.Drop, "drop", false, dropTooltip)
	flag.BoolVar(&flags.Drop, "d", false, dropTooltip)
	flag.BoolVar(&flags.Migrate, "migrate", false, migrateTooltip)
	flag.BoolVar(&flags.Migrate, "m", false, migrateTooltip)
	flag.BoolVar(&flags.MigrateDry, "migrate-dry", false, migrateDryTooltip)
	flag.BoolVar(&flags.MigrateDry, "m-dry", false, migrateDryTooltip)
	flag.BoolVar(&flags.FullReset, "full-reset", false, fullResetTooltip)
	flag.BoolVar(&flags.FullReset, "fr", false, fullResetTooltip)
	flag.BoolVar(&flags.Summary, "summary", false, summaryTooltip)
	flag.BoolVar(&flags.Summary, "s", false, summaryTooltip)

	flag.Parse()
	return flags
}
