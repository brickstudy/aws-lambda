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
			return false, fmt.Errorf("등록되지 않은 매체입니다. brickas.media 테이블을 확인해주세요.")
		}
	case "silver", "gold", "mlflow":
		return true, nil
	default:
		return false, fmt.Errorf("등록되지 않은 root 폴더입니다. 관리자에게 문의해주세요.")
	}

	// 3. check path length
	if len(parts) < 5 {
		return false, fmt.Errorf("최소 경로 depth > 5를 만족하지 못합니다. 경로를 확인해주세요.")
	}

	// 4. check project name
	if !isValidProject(projects, parts[1]) {
		return false, fmt.Errorf("등록되지 않은 프로젝트입니다. brickas.project 테이블을 확인해주세요.")
	}

	// 5. check valid date
	if !isValidDate(parts[3]) {
		return false, fmt.Errorf("잘못된 날짜 폴더입니다. 날짜 포멧을 확인해주세요. ex. yyyy-MM-dd")
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
