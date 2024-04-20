package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Email      string    `gorm:"unique"`
	Username   string
	Password   string
	ProfilePic string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Reports    []Report `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}

type Report struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()"`
	Source    string
	Amount    int
	CreatedAt time.Time
	UpdatedAt time.Time
	UserID    string // Foreign key
	Type      ReportType
}

type ReportType string

const (
	Expense ReportType = "expense"
	Income  ReportType = "income"
)
