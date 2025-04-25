package model

import "gorm.io/gorm"

// Survey model
type Survey struct {
	gorm.Model
	Title       string     `gorm:"not null"`
	Description string     `gorm:"not null"`
	Questions   []Question // A Survey has many Questions
}

// Question model
type Question struct {
	gorm.Model
	SurveyID     uint   // Foreign key to Survey
	Text         string `gorm:"not null"`            // The text of the question
	QuestionType string `gorm:"not null"`            // The type of the question (e.g., "radio", "multiple_choice", "easy_button")
	Options      string `gorm:"nullable"`            // A JSON string to store the options for the question
	Survey       Survey `gorm:"foreignKey:SurveyID"` // Foreign key relation to Survey
}

// Answer model
type Answer struct {
	gorm.Model
	UserID     uint     `gorm:"uniqueIndex:idx_user_question;not null"` // Ensure the combination of UserID and QuestionID is unique and not null
	QuestionID uint     `gorm:"uniqueIndex:idx_user_question;not null"` // Ensure the combination of UserID and QuestionID is unique and not null
	Response   string   `gorm:"not null"`                               // Ensure the response is not null
	User       User     `gorm:"foreignKey:UserID"`
	Question   Question `gorm:"foreignKey:QuestionID"`
}
