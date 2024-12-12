package controllers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/Dimasaldian/letsAdopt/app/models"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

func (server *Server) Listpet(w http.ResponseWriter, r *http.Request) {
	render := render.New(render.Options{
		Layout: "layout",
		Extensions: []string{".html"},
	})

	petModel := models.Pet{}
	pets, err := petModel.GetPets(server.DB)
	if err != nil {
		return
	}

	_ = render.HTML(w, http.StatusOK, "listpet", map[string]interface{} {
		"pets": pets,
		"showNavbar": true,
	})
}

func (server *Server) AddPet(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		// Convert age to integer
		age, err := strconv.Atoi(r.FormValue("age"))
		if err != nil {
			log.Println("Error converting age:", err)
			http.Error(w, "Invalid age format", http.StatusBadRequest)
			return
		}

		// Create the pet model instance
		pet := models.Pet{
			Name:        r.FormValue("name"),
			Type:        r.FormValue("type"),
			Age:         age,
			Description: r.FormValue("description"),
		}

		// Save the pet
		result := server.DB.Create(&pet)
		if result.Error != nil {
			http.Error(w, "Error adding pet", http.StatusInternalServerError)
			return
		}

		// After creating the pet, add the image (if any)
		imageURL := r.FormValue("image_url")
		if imageURL != "" {
			petImage := models.PetImage{
				PetID: pet.ID,
				URL:   imageURL, // Store the URL of the pet image
			}
			// Save the pet image
			result = server.DB.Create(&petImage)
			if result.Error != nil {
				log.Println("Error adding pet image:", result.Error)
			}
		}

		// Redirect after adding the pet
		http.Redirect(w, r, "/admin/pets", http.StatusFound)
		return
	}

	// If method is GET, show the form to add pet
	render := render.New(render.Options{Layout: "layout", Extensions: []string{".html"}})
	_ = render.HTML(w, http.StatusOK, "add_pet", nil)
}

func (server *Server) EditPet(w http.ResponseWriter, r *http.Request) {
	petID := mux.Vars(r)["id"]
	var pet models.Pet

	// If the method is POST, update the pet details
	if r.Method == "POST" {
		age, err := strconv.Atoi(r.FormValue("age")) // Convert age from string to int
		if err != nil {
			log.Println("Error converting age:", err)
			http.Error(w, "Invalid age format", http.StatusBadRequest)
			return
		}

		// Update pet fields
		pet.Name = r.FormValue("name")
		pet.Type = r.FormValue("type")
		pet.Age = age
		pet.Description = r.FormValue("description")

		// Save the pet data (update)
		result := server.DB.Save(&pet)
		if result.Error != nil {
			http.Error(w, "Error updating pet", http.StatusInternalServerError)
			return
		}

		// Update pet image if needed
		imageURL := r.FormValue("image_url")
		if imageURL != "" {
			// Add or update the pet image
			var petImage models.PetImage
			result = server.DB.Where("pet_id = ?", pet.ID).First(&petImage)
			if result.Error != nil {
				// If no existing image, create a new one
				petImage = models.PetImage{
					PetID: pet.ID,
					URL:   imageURL,
				}
				server.DB.Create(&petImage)
			} else {
				// If the image exists, update it
				petImage.URL = imageURL
				server.DB.Save(&petImage)
			}
		}

		// Redirect to the pet list page
		http.Redirect(w, r, "/admin/pets", http.StatusFound)
		return
	}

	// Retrieve the pet to edit
	result := server.DB.First(&pet, petID)
	if result.Error != nil {
		http.Error(w, "Pet not found", http.StatusNotFound)
		return
	}

	// Render the Edit Pet Form
	render := render.New(render.Options{Layout: "layout", Extensions: []string{".html"}})
	_ = render.HTML(w, http.StatusOK, "edit_pet", map[string]interface{}{"pet": pet})
}