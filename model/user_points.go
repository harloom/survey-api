package model

import "gorm.io/gorm"

// Define the UserPoints model
type UserPoints struct {
	gorm.Model
	UserID uint
	Points int
	User   User `gorm:"foreignKey:UserID"`
}
