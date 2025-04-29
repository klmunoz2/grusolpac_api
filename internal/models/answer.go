package models

import "gorm.io/gorm"

type Answer struct {
	gorm.Model
	QuestionID uint   `gorm:"not null;index"`
	AnswerText string `gorm:"type:text;not null"`
	IsCorrect  bool   `gorm:"not null;default:false"`
}
