package models

import (
	"time"

	"gorm.io/gorm"
)

type Enrollment struct {
	gorm.Model
	StudentID               uint      `gorm:"not null;index"`
	CourseID                uint      `gorm:"not null;index"`
	EnrollmentDate          time.Time `gorm:"default:now()"`
	TheoreticalStatus       string    `gorm:"type:varchar(50);not null;default:'pending'"`
	EvaluationScore         *float64  `gorm:"type:numeric(5,2)"`
	EvaluationDate          *time.Time
	PracticalStatus         string `gorm:"type:varchar(50);not null;default:'pending'"`
	PracticalCompletionDate *time.Time
	CertificateID           *uint
}
