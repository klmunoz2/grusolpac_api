package models

import "gorm.io/gorm"

type Module struct {
	gorm.Model
	CourseID   uint   `gorm:"not null;index"`
	Name       string `gorm:"type:varchar(255);not null"`
	OrderIndex int    `gorm:"not null;default:0"`
	Topics     []Topic
}
