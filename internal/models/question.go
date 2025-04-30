package models

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	TopicID      *uint                `gorm:"index"`
	QuestionText string               `gorm:"type:text;not null"`
	Answers      []Answer             `gorm:"foreignKey:QuestionID"`
	Evaluations  []EvaluationQuestion `gorm:"foreignKey:QuestionID"`
}
