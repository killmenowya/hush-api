package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/killmenowya/hush-api/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	// passess control to hush-routes.go
	routes.RegisterHushRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
