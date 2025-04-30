package models

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	UserID      *uint `gorm:"uniqueIndex"`
	User        *User `gorm:"foreignKey:UserID"`
	Enrollments []Enrollment
	StudentCode string `gorm:"type:varchar(50);uniqueIndex"`
}
