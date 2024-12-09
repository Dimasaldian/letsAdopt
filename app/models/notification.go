package models

import (
	"time"
)

// Notification represents the notifications table
type Notification struct {
	ID        string      `gorm:"primaryKey"`
	UserID    string     `gorm:"not null"`
	AdoptionID uint      `gorm:"not null"`
	Message   string    `gorm:"type:text;not null"`
	SentAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`

	// Relationships
	User    User    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Adoption Adoption `gorm:"foreignKey:AdoptionID;constraint:OnDelete:CASCADE"` 
}