// package fakers

// import (
// 	"math/rand"
// 	"time"

// 	"github.com/Dimasaldian/letsAdopt/app/models"
// 	"github.com/bxcodec/faker/v3"
// 	"github.com/google/uuid"
// 	"gorm.io/gorm"
// )

// func Role() string {
// 	rand.Seed(time.Now().UnixNano()) // Inisialisasi seed
// 	roles := []string{"admin", "user"}
// 	return roles[rand.Intn(len(roles))]
// }

// func UserFaker(*gorm.DB) *models.User {
// 	return &models.User{
// 		ID:        uuid.New().String(),
// 		Name:      faker.Name(),
// 		Email:     faker.Email(),
// 		Password:  "aoiwhrhq2i9irhpekg90q3hy",
// 		Role:      Role(),
// 		CreatedAt: time.Time{},
// 		UpdatedAt: time.Time{},
// 	}
// }

package fakers

import (
	"math/rand"
	"time"

	"github.com/Dimasaldian/letsAdopt/app/models"
	"github.com/bxcodec/faker/v3"
)

func Role() string {
	rand.Seed(time.Now().UnixNano()) // Inisialisasi seed
	roles := []string{"admin", "user"}
	return roles[rand.Intn(len(roles))]
}

func UserFaker() *models.User {
	return &models.User{
		Name:     faker.Name(),
		Email:    faker.Email(),
		Password: "aoiwhrhq2i9irhpekg90q3hy", // Ganti dengan password yang di-hash!
		Role:     Role(),
	}
}