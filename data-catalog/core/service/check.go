package service

import (
	"fmt"
	"strings"
	"time"

	"github.com/robert-min/aws-lambda/data-catalog/core/domain"
	"github.com/robert-min/aws-lambda/data-catalog/core/port"
)

type CheckService struct {
	projectRepo port.ProjectRepository
	mediaRepo   port.MediaRepository
}

func NewCheckService(projectRepo port.ProjectRepository, mediaRepo port.MediaRepository) *CheckService {
	return &CheckService{
		projectRepo: projectRepo,
		mediaRepo:   mediaRepo,
	}
}

func (cs CheckService) CompareNameRule(path domain.S3Path) (bool, error) {
	projects, err := cs.projectRepo.GetListUsers()
	if err != nil {
		return false, err
	}

	medias, err := cs.mediaRepo.GetListMedias()
	if err != nil {
		return false, err
	}

	// Check Name Rule
	// bronze/travel/newsapi/2024-08-05/headline_kr.json
	parts := strings.Split(path.Path, "/")

	// 1. check layer
	switch parts[0] {
	case "bronze": // silver, gold 규칙 정해지면 수정
		// 2. check source and category
		if !isValidMedia(medias, parts[2], parts[len(parts)-1]) {
			return false, fmt.Errorf("Not match brickas.media")
		}
	case "silver", "gold", "mlflow":
		return true, nil
	default:
		return false, fmt.Errorf("Not match layer")
	}

	// 3. check path length
	if len(parts) < 5 {
		return false, fmt.Errorf("Invalid path structure. lenth < 8")
	}

	// 4. check project name
	if !isValidProject(projects, parts[1]) {
		return false, fmt.Errorf("Not match brickas.project")
	}

	// 5. check valid date
	if !isValidDate(parts[3]) {
		return false, fmt.Errorf("Not match dateStr format. yyyy-MM-dd")
	}

	return true, nil
}

func isValidProject(project []domain.Project, projectName string) bool {
	for _, p := range project {
		if p.Name_ == projectName {
			return true
		}
	}
	return false
}

func isValidMedia(medias []domain.Media, mediaSource string, filename string) bool {
	var mediaCategory string
	sub_parts := strings.Split(filename, "_")
	if len(sub_parts) < 2 {
		sub_parts = strings.Split(filename, ".")
	}
	mediaCategory = sub_parts[0]

	for _, m := range medias {
		if m.Source_ == mediaSource && m.Category == mediaCategory {
			return true
		}
	}
	return false
}

func isValidDate(dateStr string) bool {
	_, err := time.Parse("2006-01-02", dateStr)
	return err == nil
}
