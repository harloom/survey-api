package model

import "gorm.io/gorm"

type PointHistory struct {
	gorm.Model
	UserID          uint
	PointsChange    int
	Description     string
	TransactionDate string `gorm:"default:CURRENT_TIMESTAMP"`
	User            User   `gorm:"foreignKey:UserID"`
}
