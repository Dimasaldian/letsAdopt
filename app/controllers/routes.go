package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (server *Server) initializeRoutes() {
	server.Router = mux.NewRouter()
	// server.Router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", http.FileServer(http.Dir("app/assets"))))

	server.Router.HandleFunc("/", server.Home).Methods("GET")
	server.Router.HandleFunc("/listpet", server.Listpet).Methods("GET")
	server.Router.HandleFunc("/dashboard", server.Dashboard).Methods("GET")

	staticFileDirectory := http.Dir("./assets/")
	staticFileHandler := http.StripPrefix("/public/", http.FileServer(staticFileDirectory))
	server.Router.PathPrefix("/public/").Handler(staticFileHandler).Methods("GET")
}