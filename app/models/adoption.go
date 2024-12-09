package models

import (
	"time"
)

// Adoption represents the adoptions table
type Adoption struct {
	ID             string      `gorm:"primaryKey"`
	UserID         string      `gorm:"not null"`
	PetID          string     `gorm:"not null"`
	Reason         string    `gorm:"type:text;not null"`
	Status         string    `gorm:"size:10;default:pending;check:status IN ('pending', 'approved', 'rejected')"`
	NotificationSent bool      `gorm:"default:false"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`

	// Relationships
	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Pet  Pet  `gorm:"foreignKey:PetID;constraint:OnDelete:CASCADE"`
}