package models

type Model struct {
	Model interface{}
}

func RegisterModels() []Model {
	return []Model{
		{Model: User{}},
		{Model: Address{}},
		{Model: Adoption{}},
		{Model: Pet{}},
		{Model: Notification{}},
		{Model: Admin{}},
	}
}
