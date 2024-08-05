package port

import "github.com/robert-min/aws-lambda/data-catalog/core/domain"

// MediaRepository is an interface for interacting with media-related data
type MediaRepository interface {
	// GetListMedias select * from brickas.media
	GetListMedias() ([]domain.Media, error)
}

// ProjectRepository is an interface for interacting with project-related data
type ProjectRepository interface {
	// GetListUsers select * from brickas.project
	GetListUsers() ([]domain.Project, error)
}

type CheckService interface {
	CompareNameRule() (bool, error)
}
