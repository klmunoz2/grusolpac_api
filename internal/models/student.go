package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	FirstName   string `gorm:"type:varchar(100);not null"`
	LastName    string `gorm:"type:varchar(100);not null"`
	Email       string `gorm:"type:varchar(255);not null;uniqueIndex"`
	Enrollments []Enrollment
}
