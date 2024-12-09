package models

// Admin represents the admins table
type Admin struct {
	ID         string   `gorm:"primaryKey"`
	UserID     string   `gorm:"not null"`
	Privileges string `gorm:"type:text"`

	// Relationships
	User User `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`
}