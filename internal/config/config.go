package config

import (
	"os"
	"strconv"
)

// Config holds all application configuration.
type Config struct {
	Port      string
	DBPath    string
	JWTSecret string
}

// Load reads configuration from environment variables with sensible defaults.
func Load() *Config {
	return &Config{
		Port:      getEnv("PANEL_PORT", "8080"),
		DBPath:    getEnv("PANEL_DB_PATH", "data/panel.db"),
		JWTSecret: getEnv("PANEL_JWT_SECRET", "change-me-in-production"),
	}
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

func getEnvInt(key string, fallback int) int {
	if v := os.Getenv(key); v != "" {
		if i, err := strconv.Atoi(v); err == nil {
			return i
		}
	}
	return fallback
}
