package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string   `gorm:"type:varchar(50);uniqueIndex;not null"`
	Email     string   `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password  string   `gorm:"type:varchar(255);not null"`
	FirstName string   `gorm:"type:varchar(100);not null"`
	LastName  string   `gorm:"type:varchar(100);not null"`
	IsActive  bool     `gorm:"default:true"`
	RoleID    uint     `gorm:"not null;index"`
	Role      Role     `gorm:"foreignKey:RoleID"`
	Student   *Student `gorm:"foreignKey:UserID"`
}
