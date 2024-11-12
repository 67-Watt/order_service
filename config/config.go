package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	_ "github.com/joho/godotenv"
)

type Config struct {
	ServerAddress string
	DatabaseDSN   string
}

// LoadConfig loads configuration from environment variables.
func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found or failed to load .env file; using system environment variables")
	}

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		getEnv("DB_USER", "root"),
		getEnv("DB_PASSWORD", "password"),
		getEnv("DB_HOST", "localhost:5432"),
		getEnv("DB_NAME", "db_order"),
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
