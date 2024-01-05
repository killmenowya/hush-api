package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/killmenowya/hush-api/pkg/models"
	"github.com/killmenowya/hush-api/pkg/utils"
)

// CREATE
func CreateTags(w http.ResponseWriter, r *http.Request) {
	// read input body and parse it for use
	CreateTag := &models.Tag{}
	utils.ParseBody(r, CreateTag)
	// call on models to access database
	t := CreateTag.CreateTags()
	res, _ := json.Marshal(t)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// READ
func GetTags(w http.ResponseWriter, r *http.Request) {
	// call on models to read data
	// business logic of sorting handled by models
	tags := models.GetTags()
	res, _ := json.Marshal(tags)
	w.Header().Set("Contrent-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// POST
func CountTagged(w http.ResponseWriter, r *http.Request) {
	// getting tag[string] as param from /tag/{tag}
	vars := mux.Vars(r)
	Tag := vars["tag"]
	// get update count
	count := models.Counter(Tag)
	// get row and database address
	tag, db := models.GetTagbyName(Tag)
	// update new count
	db.Model(&tag).Update("Counter", count)

	res, _ := json.Marshal(tag)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
