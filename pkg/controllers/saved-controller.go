package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/killmenowya/hush-api/pkg/models"
	"github.com/killmenowya/hush-api/pkg/utils"
)

// CREATE
func SaveThread(w http.ResponseWriter, r *http.Request) {
	// read input body and parse it for use
	savedthread := &models.Saved{}
	utils.ParseBody(r, savedthread)
	// call on models to access database
	s := savedthread.SaveThread()
	res, _ := json.Marshal(s)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

// DELETE
func DeleteSaved(w http.ResponseWriter, r *http.Request) {
	// getting id param from /saved/{id_saved}
	vars := mux.Vars(r)
	savedid := vars["id_saved"]
	ID, err := strconv.ParseInt(savedid, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	// call on models to access database
	saved := models.DeleteSaved(ID)
	res, _ := json.Marshal(saved)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

// READ
func GetSaved(w http.ResponseWriter, r *http.Request) {
	// call on models to read data
	saved := models.GetSaved(r)
	res, _ := json.Marshal(saved)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
