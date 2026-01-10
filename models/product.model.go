package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	UserID      string `gorm:"not null"`
	Title       string `gorm:"not null"`
	Description string `gorm:"not null"`
	Price       string `gorm:"not null"`
	Path        string `gorm:"not null"`
}
