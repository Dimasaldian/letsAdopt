package controllers

import (
	"net/http"

	"github.com/Dimasaldian/letsAdopt/app/models"
	"github.com/unrolled/render"
)

func (server *Server) Dashboard(w http.ResponseWriter, r *http.Request) {
	render := render.New(render.Options{
		Layout: "layout",
		Extensions: []string{".html"},
	})

	adminModel := models.Admin{}
	admins, err := adminModel.GetAdmin(server.DB)
	if err != nil {
		return
	}

	_ = render.HTML(w, http.StatusOK, "dashboard", map[string]interface{} {
		"title": "Dashboard",
		"showNavbar": false,
		"admin": admins,
	})
}