package main

import (
	"log/slog"

	"github.com/robert-min/aws-lambda/data-catalog/adapter/config"
	"github.com/robert-min/aws-lambda/data-catalog/adapter/storage/mysql"
	"github.com/robert-min/aws-lambda/data-catalog/adapter/storage/mysql/repository"
	"github.com/robert-min/aws-lambda/data-catalog/core/domain"
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

	path := domain.S3Path{
		Bucket: "brickstudy",
		Path:   "s3://brickstudy/bronze/travel/newsapi/2024-08-05/headline_kr.json",
	}

	_, err = service.CompareNameRule(path)
	if err != nil {
		slog.Error("Error to compare name rule.", "error", err)
	}
}
