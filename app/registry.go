package app

import "github.com/Dimasaldian/letsAdopt/app/models"

type Model struct {
	Model interface{}
}

func RegisterModels() []Model {
	return []Model{
		{Model: models.User{}},  
		{Model: models.Address{}}, 
		{Model: models.Adoption{}},
		{Model: models.Pet{}},  
		{Model: models.Notification{}},
		{Model: models.Admin{}}, 
	}
}