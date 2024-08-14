package repository

import (
	"github.com/robert-min/aws-lambda/data-catalog/adapter/storage/mysql"
	"github.com/robert-min/aws-lambda/data-catalog/core/domain"
)

type MediaRepository struct {
	db *mysql.DB
}

func NewMediaRepository(db *mysql.DB) *MediaRepository {
	return &MediaRepository{
		db: db,
	}
}

func (mr *MediaRepository) GetListMedias() ([]domain.Media, error) {
	query := "SELECT * FROM media"
	rows, err := mr.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var medias []domain.Media
	for rows.Next() {
		var media domain.Media
		err := rows.Scan(&media.Idx, &media.Source_, &media.Category, &media.Url, &media.ClientID, &media.ClientPW, &media.Token)
		if err != nil {
			return nil, err
		}

		medias = append(medias, media)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return medias, nil
}
