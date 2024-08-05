package main

import (
	"fmt"
	"log/slog"

	"github.com/robert-min/aws-lambda/data-catalog/adapter/config"
	"github.com/robert-min/aws-lambda/data-catalog/adapter/storage/mysql"
	"github.com/robert-min/aws-lambda/data-catalog/adapter/storage/mysql/repository"
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

	projectRepo := repository.NewProjectRepository(db)
	_, err = projectRepo.GetListUsers()
	if err != nil {
		slog.Error("Error to get project list.", "error", err)
	}
	// fmt.Println(result)

	mediaRepo := repository.NewMediaRepository(db)
	result, err := mediaRepo.GetListMedias()
	if err != nil {
		slog.Error("Error to get media list.", "error", err)
	}
	fmt.Println(result)
}
