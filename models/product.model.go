package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	UserID      string `gorm:"not null"`
	Title       string `gorm:"size:20;not null"`
	Description string `gorm:"size:255;not null"`
	Path        string `gorm:"not null"`
}
