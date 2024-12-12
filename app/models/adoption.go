// package models

// import (
// 	"time"
// )

// // Adoption represents the adoptions table
// type Adoption struct {
// 	ID             string      `gorm:"primaryKey"`
// 	UserID         string      `gorm:"not null"`
// 	PetID          string     `gorm:"not null"`
// 	Reason         string    `gorm:"type:text;not null"`
// 	Status         string    `gorm:"size:10;default:pending;check:status IN ('pending', 'approved', 'rejected')"`
// 	NotificationSent bool      `gorm:"default:false"`
// 	CreatedAt      time.Time `gorm:"autoCreateTime"`
// 	UpdatedAt      time.Time `gorm:"autoUpdateTime"`

// 	// Relationships
// 	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
// 	Pet  Pet  `gorm:"foreignKey:PetID;constraint:OnDelete:CASCADE"`
// }

package models

import (
	"gorm.io/gorm"
)

type Adoption struct {
	gorm.Model
	IDAdopt        uint   `gorm:"primaryKey;autoIncrement:true" json:"id_adopt"`
	UserID         uint   `json:"user_id"`
	User           User   `gorm:"foreignKey:UserID;references:ID" json:"user"`
	PetID          uint   `json:"pet_id"`
	Pet            Pet    `gorm:"foreignKey:PetID;references:ID" json:"pet"`
	Reason         string `gorm:"type:text;not null" json:"reason"`
	Status         string `gorm:"size:10;not null;default:'pending'" json:"status"`
	NotificationSent bool `json:"notification_sent"`
	// Notifications []Notification `gorm:"foreignKey:AdoptionID;references:IDAdopt" json:"notifications"`
}