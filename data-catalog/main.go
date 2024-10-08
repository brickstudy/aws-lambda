package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/robert-min/aws-lambda/data-catalog/adapter/config"
	"github.com/robert-min/aws-lambda/data-catalog/adapter/storage/mysql"
	"github.com/robert-min/aws-lambda/data-catalog/adapter/storage/mysql/repository"
	"github.com/robert-min/aws-lambda/data-catalog/core/domain"
	"github.com/robert-min/aws-lambda/data-catalog/core/service"
)

func process(bucket string, object string) {
	config, err := config.New()
	if err != nil {
		log.Printf("Error to set config. : %v", err)
	}
	log.Println("Config loaded successfully")

	db, err := mysql.New(config.DB)
	if err != nil {
		log.Printf("Error to connect database. : %v", err)
	}
	defer db.Close()
	log.Println("Database successfully")

	// Dependency injection
	// Check
	projectRepo := repository.NewProjectRepository(db)
	mediaRepo := repository.NewMediaRepository(db)
	serviceCheck := service.NewCheckService(projectRepo, mediaRepo)

	path := domain.S3Path{
		Bucket: bucket,
		Path:   object,
	}

	flag, err := serviceCheck.CompareNameRule(path)
	if err != nil {
		log.Printf("Error to compare name rule. : %v", err)
	}
	log.Println("Compare successfully")

	// send message
	err = service.SendDiscordMessage(flag, object, err)
	if err != nil {
		log.Printf("Error to send message: %v", err)
	}
	log.Println("send message")
}

func HandleRequest(ctx context.Context, s3Event events.S3Event) {
	for _, record := range s3Event.Records {
		s3 := record.S3
		bucket := s3.Bucket.Name
		object := s3.Object.Key

		// S3 이벤트로부터 파일 경로를 로그에 출력
		log.Printf("Bucket: %s, Object: %s", bucket, object)

		// 추가적인 처리 로직을 여기에 추가
		process(bucket, object)
	}
}

func main() {
	// _, err := config.New()
	// if err != nil {
	// 	log.Printf("Error to set config. : %v", err)
	// }
	// log.Println("Config loaded successfully")
	// service.SendDiscordMessage(true, "hello", nil)
	lambda.Start(HandleRequest)
}
