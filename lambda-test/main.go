package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, s3Event events.S3Event) {
	for _, record := range s3Event.Records {
		s3 := record.S3
		bucket := s3.Bucket.Name
		object := s3.Object.Key

		// S3 이벤트로부터 파일 경로를 로그에 출력
		log.Printf("Bucket: %s, Object: %s", bucket, object)

		// 추가적인 처리 로직을 여기에 추가
	}
}

func main() {
	lambda.Start(HandleRequest)
}
