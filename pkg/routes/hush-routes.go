package routes

import (
	"github.com/gorilla/mux"
	"github.com/killmenowya/hush-api/pkg/controllers"
)

var RegisterHushRoutes = func(router *mux.Router) {
	// list of route for various functions

	// comment
	router.HandleFunc("/comment/", controllers.CreateComment).Methods("POST")
	router.HandleFunc("/comment/{id_thread}", controllers.GetComment).Methods("GET")
	router.HandleFunc("/comment/{id_comment}", controllers.DeleteComment).Methods("DELETE")
	router.HandleFunc("/comment/{id_comment}", controllers.EditComment).Methods("PUT")

	// thread
	router.HandleFunc("/thread/", controllers.CreateThread).Methods("POST")
	router.HandleFunc("/thread/", controllers.GetThread).Methods("GET")
	router.HandleFunc("/thread/{id_thread}", controllers.DeleteThread).Methods("DELETE")
	router.HandleFunc("/thread/{id_thread}", controllers.EditThread).Methods("PUT")
	router.HandleFunc("/thread/user/{id_user}", controllers.GetThreadByUser).Methods("GET")
	router.HandleFunc("/thread/tag/{tag}", controllers.GetThreadByTag).Methods("GET")
	router.HandleFunc("/thread/{id_thread}", controllers.GetThreadByID).Methods("GET")
	router.HandleFunc("/thread/comment/{id_thread}", controllers.CountResponse).Methods("PUT")

	// tags
	router.HandleFunc("/tag/", controllers.GetTags).Methods("GET")
	router.HandleFunc("/tag/", controllers.CreateTags).Methods("POST")
	router.HandleFunc("/tag/{tag}", controllers.CountTagged).Methods("PUT")

	// saved
	router.HandleFunc("/saved/", controllers.SaveThread).Methods("POST")
	router.HandleFunc("/saved/{id_saved}", controllers.DeleteSaved).Methods("DELETE")
	router.HandleFunc("/saved/", controllers.GetSaved).Methods("GET")

}
