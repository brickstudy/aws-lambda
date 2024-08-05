package service

import (
	"fmt"

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

func (cs CheckService) CompareNameRule() (bool, error) {
	projects, err := cs.projectRepo.GetListUsers()
	if err != nil {
		return false, err
	}
	fmt.Println(projects)

	medias, err := cs.mediaRepo.GetListMedias()
	if err != nil {
		return false, err
	}
	fmt.Println(medias)
	return true, nil
}
