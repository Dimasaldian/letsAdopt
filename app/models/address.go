// package models

// import (
// 	"time"
// )

// // Address represents the addresses table
// type Address struct {
// 	ID        string      `gorm:"primaryKey"`
// 	UserID    *string     `gorm:"index"`
// 	PetID     *string     `gorm:"index"`
// 	Street    string    `gorm:"size:100"`
// 	City      string    `gorm:"size:50"`
// 	State     string    `gorm:"size:50"`
// 	PostalCode string    `gorm:"size:20"`
// 	CreatedAt time.Time `gorm:"autoCreateTime"`
// 	UpdatedAt time.Time `gorm:"autoUpdateTime"`

// 	// Relationships
// 	User *User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
// 	Pet  *Pet  `gorm:"foreignKey:PetID;constraint:OnDelete:CASCADE"`
// }

package models

import (
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model
	UserID      uint   `json:"user_id"`
	User        User   `gorm:"foreignKey:UserID;references:ID" json:"user"`
	PetID       uint   `json:"pet_id"`
	Pet         Pet    `gorm:"foreignKey:PetID;references:ID" json:"pet"`
	Street      string `gorm:"size:100" json:"street"`
	City        string `gorm:"size:50" json:"city"`
	State       string `gorm:"size:50" json:"state"`
	PostalCode string `gorm:"size:20" json:"postal_code"`
}