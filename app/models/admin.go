package models

import (
	"gorm.io/gorm"
)

type Admin struct {
	gorm.Model
	AdminID   uint   `gorm:"primaryKey;autoIncrement:true" json:"admin_id"`
	UserID    uint   `json:"user_id"`
	User      User   `gorm:"foreignKey:UserID;references:ID" json:"user"`
	Privileges string `gorm:"type:text" json:"privileges"`
}

func (a *Admin) GetAdmin(db *gorm.DB) (*[]Admin, error) {
	var err error
	var admins []Admin

	err = db.Debug().Model(&Admin{}).Preload("User").Limit(20).Find(&admins).Error
	if err != nil {
		return nil, err
	}

	return &admins, nil
}