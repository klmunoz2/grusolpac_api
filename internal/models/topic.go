package models

import "gorm.io/gorm"

type Topic struct {
	gorm.Model
	ModuleID   uint   `gorm:"not null;index"`
	Name       string `gorm:"type:varchar(255);not null"`
	Content    string `gorm:"type:text"`
	OrderIndex int    `gorm:"not null;default:0"`
	Questions  []Question
}
