package models

import (
	"time"
)

type Pet struct {
	ID          string      `gorm:"primaryKey"`
	Name        string    `gorm:"size:100;not null"`
	Type        string    `gorm:"size:20;not null;check:type IN ('dog', 'cat', 'bird', 'other')"`
	Breed       string    `gorm:"size:100"`
	Age         int
	Description string
	Vaccinated  bool      `gorm:"default:false"`
	Address     Address    `gorm:"size:255"` 
	CreatedAt   time.Time `gorm:"autoCreateTime"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime"`
}