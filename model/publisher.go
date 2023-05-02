package model

import "gorm.io/gorm"

// Publisher model
type Publisher struct {
	gorm.Model

	Publisher string `gorm:"not null" json:"publisher"`
	Country   string `gorm:"not null" json:"country"`
}
