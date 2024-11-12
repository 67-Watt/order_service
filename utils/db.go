package utils

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"order_service/config"
)

// ConnectDB establishes a connection to the PostgreSQL database using GORM.
func ConnectDB(cfg *config.Config) (*gorm.DB, error) {
	dsn := cfg.DatabaseDSN
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Connected to PostgreSQL database using GORM successfully.")
	return db, nil
}
