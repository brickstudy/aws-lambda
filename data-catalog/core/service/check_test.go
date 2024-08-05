package service_test

import (
	"fmt"
	"testing"

	"github.com/robert-min/aws-lambda/data-catalog/core/domain"
	"github.com/robert-min/aws-lambda/data-catalog/core/service"
)

type MockProjectRepo struct{}
type MockMediaRepo struct{}

func (mpr MockProjectRepo) GetListUsers() ([]domain.Project, error) {
	return []domain.Project{
		{Idx: 1, Name_: "project1", Admin: "admin1"},
		{Idx: 2, Name_: "project2", Admin: "admin2"},
		{Idx: 3, Name_: "project3", Admin: "admin3"},
	}, nil
}

func (mmr MockMediaRepo) GetListMedias() ([]domain.Media, error) {
	return []domain.Media{
		{Idx: 1, Source_: "source1", Category: "headline", Url: "url1"},
		{Idx: 2, Source_: "source2", Category: "news", Url: "url2"},
	}, nil
}

func TestCompareNameRule(t *testing.T) {
	mockProjectRepo := MockProjectRepo{}
	mockMediaRepo := MockMediaRepo{}

	cs := service.NewCheckService(mockProjectRepo, mockMediaRepo)

	tests := []struct {
		s3Path    domain.S3Path
		wantValid bool
		wantErr   error
	}{
		{
			s3Path:    domain.S3Path{Bucket: "testbucket", Path: "s3://brickstudy/bronze/project1/source1/2024-08-05/headline_kr.json"},
			wantValid: true,
			wantErr:   nil,
		},
		{
			s3Path:    domain.S3Path{Bucket: "testbucket", Path: "s3://brickstudy/silver/project1/source1/2024-08-05/headline_kr.json"},
			wantValid: true,
			wantErr:   nil,
		},
		{
			s3Path:    domain.S3Path{Bucket: "testbucket", Path: "s3://brickstudy/bronze/source2/project2/2024-08-05/news_kr.json"},
			wantValid: false,
			wantErr:   fmt.Errorf("Not match brickas.media"),
		},
		// {
		// 	s3Path:    domain.S3Path{Bucket: "testbucket", Path: "s3://brickstudy/mlflow/source1/project1/2024-08-05/headline_kr.json"},
		// 	wantValid: true,
		// 	wantErr:   false,
		// },
		// {
		// 	s3Path:    domain.S3Path{Bucket: "testbucket", Path: "s3://brickstudy/invalid/source1/project1/2024-08-05/headline_kr.json"},
		// 	wantValid: false,
		// 	wantErr:   true,
		// },
		// {
		// 	s3Path:    domain.S3Path{Bucket: "testbucket", Path: "s3://brickstudy/bronze/source1/project1/invalid-date/headline_kr.json"},
		// 	wantValid: false,
		// 	wantErr:   true,
		// },
		// {
		// 	s3Path:    domain.S3Path{Bucket: "testbucket", Path: "s3://brickstudy/bronze/source1/invalidproject/2024-08-05/headline_kr.json"},
		// 	wantValid: false,
		// 	wantErr:   true,
		// },
		// {
		// 	s3Path:    domain.S3Path{Bucket: "testbucket", Path: "s3://brickstudy/bronze/source1/project1/2024-08-05/invalidcategory_kr.json"},
		// 	wantValid: false,
		// 	wantErr:   true,
		// },
		// {
		// 	s3Path:    domain.S3Path{Bucket: "testbucket", Path: "s3://brickstudy/bronze/source1/project1/2024-08-05"},
		// 	wantValid: false,
		// 	wantErr:   true,
		// },
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Path: %s", tt.s3Path.Path), func(t *testing.T) {
			gotValid, err := cs.CompareNameRule(tt.s3Path)

			if gotValid != tt.wantValid {
				t.Errorf("CompareNameRule() =: %v, want: %v", gotValid, tt.wantValid)
			}
			if err != tt.wantErr {
				t.Errorf("Error: %v, want: %v", err, tt.wantErr)
			}

		})
	}
}
