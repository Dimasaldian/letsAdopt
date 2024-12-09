package models

import (
	"time"
)

// Address represents the addresses table
type Address struct {
	ID        string      `gorm:"primaryKey"`
	UserID    *string     `gorm:"index"`
	PetID     *string     `gorm:"index"`
	Street    string    `gorm:"size:100"`
	City      string    `gorm:"size:50"`
	State     string    `gorm:"size:50"`
	PostalCode string    `gorm:"size:20"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`

	// Relationships
	User *User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
	Pet  *Pet  `gorm:"foreignKey:PetID;constraint:OnDelete:CASCADE"`
}