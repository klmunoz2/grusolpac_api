package models

import "gorm.io/gorm"

type EvaluationQuestion struct {
	gorm.Model
	EvaluationID uint             `gorm:"primaryKey"`
	QuestionID   uint             `gorm:"primaryKey"`
	Evaluation   CourseEvaluation `gorm:"foreignKey:EvaluationID;references:ID"`
	Question     Question         `gorm:"foreignKey:QuestionID;references:ID"`
}
