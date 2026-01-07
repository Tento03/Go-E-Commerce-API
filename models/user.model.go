package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserID   string `gorm:"not null;unique"`
	Username string `gorm:"not null;unique"`
	Password string `gorm:"not null"`
}
