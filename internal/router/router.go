package router

import (
	"log"

	"github.com/Isaac799/postgres-sql-migrations/internal/config"
	"github.com/Isaac799/postgres-sql-migrations/internal/handlers"
)

func Router() {
	flags := parseCliFlags()

	if flags.Help {
		handlers.ShowHelp()
		return
	}

	if flags.GenMigration != "" {
		if err := handlers.GenerateMigration(flags.GenMigration); err != nil {
			log.Fatalf("Error generating migration: %v", err)
		}
		return
	}

	cfg, err := config.LoadConfig(flags.Env)
	if err != nil {
		log.Fatalf("%v\nConsider the example.env.bash\nMaybe you forgot '-e dev'?", err)
	}

	if flags.ListMigrations {
		if err := handlers.ListAppliedMigrations(cfg); err != nil {
			log.Fatalf("Error listing migrations: %v", err)
		}
		return
	}

	if flags.Summary {
		if err := handlers.SummarizeMigrations(cfg); err != nil {
			log.Fatalf("Error summarizing migrations: %v", err)
		}
		return
	}

	if flags.MigrateDry {
		if err := handlers.MigrateDryRun(cfg); err != nil {
			log.Fatalf("Error migrating: %v", err)
		}
		return
	}

	if flags.Create {
		if err := handlers.CreateDatabase(cfg); err != nil {
			log.Fatalf("Error creating database: %v", err)
		}
		return
	}

	if flags.Migrate {
		if err := handlers.Migrate(cfg); err != nil {
			log.Fatalf("Error migrating: %v", err)
		}
		return
	}

	if flags.FullReset {
		if err := handlers.FullReset(cfg, flags.Env); err != nil {
			log.Fatalf("Error performing full reset: %v", err)
		}
		return
	}

	if flags.Drop {
		if err := handlers.DropDatabase(cfg, flags.Env); err != nil {
			log.Fatalf("Error dropping database: %v", err)
		}
		return
	}
}
