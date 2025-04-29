package models

import (
	"time"

	"gorm.io/gorm"
)

// Base estructura base para todos los modelos
type Base struct {
	ID        uint           `json:"id" gorm:"primarykey"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// Ejemplo de un modelo de usuario
type User struct {
	Base
	Username  string `json:"username" gorm:"uniqueIndex;not null"`
	Email     string `json:"email" gorm:"uniqueIndex;not null"`
	Password  string `json:"-" gorm:"not null"` // No se devuelve en JSON
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	IsActive  bool   `json:"is_active" gorm:"default:true"`
}

// Puedes agregar más modelos según tus necesidades
// Ejemplo:
// type Product struct {
//     Base
//     Name        string  `json:"name" gorm:"not null"`
//     Description string  `json:"description"`
//     Price       float64 `json:"price" gorm:"not null"`
//     Stock       int     `json:"stock" gorm:"default:0"`
// }
