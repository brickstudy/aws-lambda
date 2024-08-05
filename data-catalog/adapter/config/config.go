package config

import (
	"os"

	"github.com/joho/godotenv"
)

type (
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

func New() (*Container, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
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
