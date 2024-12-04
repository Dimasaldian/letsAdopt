package controllers

import (
	"net/http"

	"github.com/unrolled/render"
)

func Listpet(w http.ResponseWriter, r *http.Request) {
	render := render.New(render.Options{
		Layout: "layout",
		Extensions: []string{".html"},
	})

	_ = render.HTML(w, http.StatusOK, "listpet", map[string]interface{} {
		"title": "Home Title",
		"body": "Home Description",
	})
}