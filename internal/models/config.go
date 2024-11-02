package models

import "fmt"

type Config struct {
	Name     string
	User     string
	Password string
	Host     string
	Port     string
	SslMode  string
}

func (cfg *Config) ConnStrDatabase() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.Name, cfg.SslMode)
}

func (cfg *Config) ConnStr() string {
	return fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=postgres sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.SslMode)
}
