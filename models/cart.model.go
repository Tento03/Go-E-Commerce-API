package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	UserID string `gorm:"not null"`
	Items  []CartItem
}

type CartItem struct {
	gorm.Model
	CartID    string
	ProductID string
	Qty       int `gorm:"not null"`
	Product   Product
}
