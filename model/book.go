package model

import "gorm.io/gorm"

// Book model
type Book struct {
	gorm.Model

	Title       string `gorm:"not null" json:"title"`
	Description string `gorm:"not null" json:"description"`
	Author      string `gorm:"not null" json:"author"`
}
