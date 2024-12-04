package fakers

import (
	"github.com/Dimasaldian/letsAdopt/app/models"
	"github.com/Dimasaldian/letsAdopt/database/fakers"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func UserFaker(*gorm.DB) *models.User {
	return &models.User{
		ID:			uuid.New().String(),
		Name: 		fakers.Name(),
		Email:		fakers.Email(),
		Password:	"aoiwhrhq2i9irhpekg90q3hy",
		Role:   	fakers.Role(),
		CreatedAt: time.Time, 
		UpdatedAt: time.Time,
	}
}