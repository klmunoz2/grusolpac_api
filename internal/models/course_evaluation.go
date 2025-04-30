package models

import "gorm.io/gorm"

type CourseEvaluation struct {
	gorm.Model
	CourseID          uint                 `gorm:"not null;uniqueIndex"`
	NumberOfQuestions int                  `gorm:"not null"`
	PassingScore      float64              `gorm:"type:numeric(5,2);not null"`
	Questions         []EvaluationQuestion `gorm:"foreignKey:EvaluationID"`
}
