// package models

// import (
// 	"time"
// )

// // Notification represents the notifications table
// type Notification struct {
// 	ID        string      `gorm:"primaryKey"`
// 	UserID    string     `gorm:"not null"`
// 	AdoptionID uint      `gorm:"not null"`
// 	Message   string    `gorm:"type:text;not null"`
// 	SentAt    time.Time `gorm:"default:CURRENT_TIMESTAMP"`

// 	// Relationships
// 	User    User    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
// 	Adoption Adoption `gorm:"foreignKey:AdoptionID;constraint:OnDelete:CASCADE"` 
// }

package models

import (
	"gorm.io/gorm"
	"time"
)

type Notification struct {
	gorm.Model
	NotificationID uint     `gorm:"primaryKey;autoIncrement:true" json:"notification_id"`
	UserID          uint     `json:"user_id"`
	User            User     `gorm:"foreignKey:UserID;references:ID" json:"user"`
	AdoptionID      uint     `json:"adoption_id"`
	Adoption        Adoption `gorm:"foreignKey:AdoptionID;references:IDAdopt" json:"adoption"`
	Message         string   `gorm:"type:text;not null" json:"message"`
	SentAt          time.Time `json:"sent_at"`
}