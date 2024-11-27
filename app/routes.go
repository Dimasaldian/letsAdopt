package app

import (
	"github.com/Dimasaldian/letsAdopt/app/controllers"
)

func (server *Server) initializeRoutes() {
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
}