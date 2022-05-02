package models

import "gorm.io/gorm"

type Link struct {
	gorm.Model
	LinkID   string `gorm:"unique, not null"`
	Original string `gorm:"not null"`
}
