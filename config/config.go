package config

import (
	"fmt"
	"os"
)

type Config struct {
	ServerAddress string
	DatabaseDSN   string
}

// LoadConfig loads configuration from environment variables.
func LoadConfig() (*Config, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "password"),
		getEnv("DB_HOST", "localhost:5432"),
		getEnv("DB_NAME", "order_service"),
	)

	return &Config{
		ServerAddress: getEnv("SERVER_ADDRESS", ":8080"),
		DatabaseDSN:   dsn,
	}, nil
}

// getEnv gets an environment variable or returns a default value.
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
