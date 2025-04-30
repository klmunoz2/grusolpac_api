package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Certificate struct {
	gorm.Model
	EnrollmentID      uint      `gorm:"not null;uniqueIndex"`
	IssueDate         time.Time `gorm:"default:now()"`
	QRCodeURL         string    `gorm:"type:text;not null"`
	PDFPath           *string   `gorm:"type:text"`
	VerificationToken uuid.UUID `gorm:"type:uuid;not null;uniqueIndex;default:gen_random_uuid()"`
}
