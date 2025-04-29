package models

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	Name               string `gorm:"type:varchar(255);not null"`
	Description        string `gorm:"type:text"`
	TotalHours         int    `gorm:"not null"`
	HasFinalEvaluation bool   `gorm:"not null;default:true"`
	Modules            []Module
	Evaluations        []CourseEvaluation
	Enrollments        []Enrollment
}
