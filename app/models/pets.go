package models

import (

	"gorm.io/gorm"
)

type Pet struct {
	gorm.Model
	Name        string    `gorm:"size:100;not null" json:"name"`
	Type        string    `gorm:"size:20;not null" json:"type"` // dog, cat, bird, other
	Breed       string    `gorm:"size:100" json:"breed"`
	Age         int       `json:"age"`
	Description string    `gorm:"type:text" json:"description"`
	Vaccinated  bool      `json:"vaccinated"`
	Images      []PetImage `gorm:"foreignKey:PetID;references:ID" json:"images"`      // One-to-many relationship with PetImage
}

func (p *Pet) GetPets(db *gorm.DB) (*[]Pet, error) {
	var err error
	var pets []Pet

	err = db.Debug().Model(&Pet{}).Limit(20).Find(&pets).Error
	if err != nil {
		return nil, err
	}

	return &pets, nil
}