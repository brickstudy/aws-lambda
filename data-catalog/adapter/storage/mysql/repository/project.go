package repository

import (
	"github.com/robert-min/aws-lambda/data-catalog/adapter/storage/mysql"
	"github.com/robert-min/aws-lambda/data-catalog/core/domain"
)

type ProjectRepository struct {
	db *mysql.DB
}

func NewProjectRepository(db *mysql.DB) *ProjectRepository {
	return &ProjectRepository{
		db,
	}
}

// Get ListUers all users from database
func (pr *ProjectRepository) GetListUsers() ([]domain.Project, error) {
	query := "SELECT * FROM project"
	rows, err := pr.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var projects []domain.Project
	for rows.Next() {
		var project domain.Project
		err := rows.Scan(&project.Idx, &project.Name_, &project.Admin)
		if err != nil {
			return nil, err
		}

		projects = append(projects, project)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}
