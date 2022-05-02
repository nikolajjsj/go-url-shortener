package database

import (
	"go-url-shortener/internal/database/models"
)

func (db *DB) CreateLink(id string, original string) (*models.Link, error) {
	link := models.Link{
		Original: original,
		LinkID:   id,
	}
	if err := db.client.Create(&link).Error; err != nil {
		return nil, err
	}
	return &link, nil
}

func (db *DB) GetLink(id string) (*models.Link, error) {
	var link models.Link
	if err := db.client.First(&link, "link_id == ?", id).Error; err != nil {
		return nil, err
	}
	return &link, nil
}
