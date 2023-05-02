package model

import (
	"gorm.io/gorm"
)

// User struct
type User struct {
	gorm.Model

	Username string `gorm:"type:varchar(512);uniqueIndex;not null" json:"username"`
	Email    string `gorm:"type:varchar(512);uniqueIndex;not null" json:"email"`
	Password string `gorm:"not null" json:"password"`
	Name     string `json:"name"`
}
