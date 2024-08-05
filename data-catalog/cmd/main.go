package main

import (
	"log/slog"

	"github.com/robert-min/aws-lambda/data-catalog/adapter/config"
	"github.com/robert-min/aws-lambda/data-catalog/adapter/storage/mysql"
	"github.com/robert-min/aws-lambda/data-catalog/adapter/storage/mysql/repository"
	"github.com/robert-min/aws-lambda/data-catalog/core/service"
)

func main() {
	config, err := config.New()
	if err != nil {
		slog.Error("Error to set config.", "error", err)
	}
	db, err := mysql.New(config.DB)
	if err != nil {
		slog.Error("Error to connect database.", "error", err)
	}
	defer db.Close()

	// Dependency injection
	// Check
	projectRepo := repository.NewProjectRepository(db)
	mediaRepo := repository.NewMediaRepository(db)
	service := service.NewCheckService(projectRepo, mediaRepo)
	_, err = service.CompareNameRule()
	if err != nil {
		slog.Error("Error to compare name rule.", "error", err)
	}
}
