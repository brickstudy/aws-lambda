package config

import (
	"os"

	"github.com/joho/godotenv"
)

type (
	// Config container
	Container struct {
		DB *DB
	}

	// Database contains all the environment variables for the database
	DB struct {
		Username string
		Password string
		Hostname string
		Port     string
		Name     string
	}
)

// Create New Config
func New() (*Container, error) {
	// Check deployment type
	deployment := os.Getenv("DEPLOYMENT")
	if deployment != "prod" {
		err := godotenv.Load()
		if err != nil {
			return nil, err
		}
	}

	db := &DB{
		Username: os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Hostname: os.Getenv("DB_HOSTNAME"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
	}

	return &Container{
		db,
	}, nil
}
