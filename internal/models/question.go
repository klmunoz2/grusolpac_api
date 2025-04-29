package models

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	TopicID      *uint  `gorm:"index"`
	QuestionText string `gorm:"type:text;not null"`
	Answers      []Answer
	Evaluations  []EvaluationQuestion
}
