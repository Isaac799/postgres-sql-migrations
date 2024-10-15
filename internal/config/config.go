package config

import (
	"fmt"
	"os"
	"postgres_sql_migrations/internal/models"
)

func loadConfig(env string) (*models.Config, error) {
	cfg := &models.Config{}
	keys := []string{"DBNAME", "DBUSER", "DBPASS", "DBHOST", "DBPORT", "DBSSL"}

	for _, key := range keys {
		value := os.Getenv(fmt.Sprintf("%s_%s", env, key))
		if value == "" {
			return nil, fmt.Errorf("missing environment variable: %s_%s", env, key)
		}

		switch key {
		case "DBNAME":
			cfg.Name = value
		case "DBUSER":
			cfg.User = value
		case "DBPASS":
			cfg.Password = value
		case "DBHOST":
			cfg.Host = value
		case "DBPORT":
			cfg.Port = value
		case "DBSSL":
			cfg.SslMode = value
		}
	}

	return cfg, nil
}

func LoadConfig(env string) (*models.Config, error) {
	cfg, err := loadConfig(env)
	if err != nil {
		return nil, fmt.Errorf("Failed loading environment '%s': %w", env, err)
	}
	return cfg, nil
}
