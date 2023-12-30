package routes

import (
	"github.com/gorilla/mux"
	"github.com/killmenowya/hush-api/pkg/controllers"
)

var RegisterHushRoutes = func(router *mux.Router) {
	router.HandleFunc("/comment/", controllers.CreateComment).Methods("POST")
}
