package models

type EvaluationQuestion struct {
	EvaluationID uint `gorm:"primaryKey"`
	QuestionID   uint `gorm:"primaryKey"`
}
