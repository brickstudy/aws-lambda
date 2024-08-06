package service_test

import (
	"fmt"
	"testing"

	"github.com/robert-min/aws-lambda/data-catalog/core/service"
)

func TestSendDiscordMessage(t *testing.T) {
	tests := []struct {
		flag  bool
		path  string
		error error
	}{
		{
			flag:  true,
			path:  "bronze/project1/source1/2024-08-05/headline_kr.json",
			error: nil,
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("Path: %s", tt.path), func(t *testing.T) {
			service.SendDiscordMessage(tt.flag, tt.path, tt.error)
		})
	}

}
