package models

import (
	"time"
)

type PetImage struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	PetID     string    `gorm:"not null;index"` // Foreign key to Pet model
	URL       string    `gorm:"size:255;not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}