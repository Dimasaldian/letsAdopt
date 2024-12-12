package controllers

import (
	"net/http"

	"github.com/Dimasaldian/letsAdopt/app/models"
	"github.com/unrolled/render"
)


func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	render := render.New(render.Options{
		Layout: "layout",
		Extensions: []string{".html"},
	})

	// petModel := models.Pet{}
	// pets, err := petModel.GetPets(server.DB) // Fungsi untuk mengambil data pet dari database
	// 	if err != nil {
	// 		return
	// 	}
	pets := []models.Pet{}
	result := server.DB.Debug().Limit(4).Find(&pets) // Aktifkan debug GORM
	if result.Error != nil {
		return
	}

	_ = render.HTML(w, http.StatusOK, "home", map[string]interface{} {
		"title": "Home Title",
		"body": "Home Description",
		"showNavbar": true,
		"pets": pets,
	})
}
