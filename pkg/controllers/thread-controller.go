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
func CreateThread(w http.ResponseWriter, r *http.Request) {
	// read input body and parse it for use
	CreateThread := &models.Thread{}
	utils.ParseBody(r, CreateThread)
	// call on models to access database
	t := CreateThread.CreateThread()
	res, _ := json.Marshal(t)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// DELETE
func DeleteThread(w http.ResponseWriter, r *http.Request) {
	// getting id as param from /thread/{id_thread}
	vars := mux.Vars(r)
	threadid := vars["id_thread"]
	ID, err := strconv.ParseInt(threadid, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	// call on models to access database
	thread := models.DeleteThread(ID)
	res, _ := json.Marshal(thread)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

// READ
func GetThread(w http.ResponseWriter, r *http.Request) {
	// call on models to read data
	thread := models.GetThread(r)
	res, _ := json.Marshal(thread)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetThreadByID(w http.ResponseWriter, r *http.Request) {
	// getting id as param from /thread/{id_thread}
	vars := mux.Vars(r)
	threadid := vars["id_thread"]
	ID, err := strconv.ParseInt(threadid, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	// call on models to access database
	thread, _ := models.GetThreadbyID(ID)
	res, _ := json.Marshal(thread)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetThreadByUser(w http.ResponseWriter, r *http.Request) {
	// getting id as param from /thread/user/{id_user}
	vars := mux.Vars(r)
	threadid := vars["id_user"]
	ID, err := strconv.ParseInt(threadid, 0, 0)
	if err != nil {
		fmt.Println("erro while parsing")
	}
	// call on models to access database
	page := models.GetThreadbyUser(r, ID)
	res, _ := json.Marshal(page)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetThreadByTag(w http.ResponseWriter, r *http.Request) {
	// getting tag[string] as param from /thread/{tag}
	vars := mux.Vars(r)
	tag := vars["tag"]
	// call on models to access database
	page := models.GetThreadbyTag(r, tag)
	res, _ := json.Marshal(page)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// UPDATE
func EditThread(w http.ResponseWriter, r *http.Request) {
	// getting the id as param from /thread/{id_thread}
	vars := mux.Vars(r)
	threadid := vars["id_thread"]
	ID, err := strconv.ParseInt(threadid, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	// getting the original thread and db address
	ThreadDetails, db := models.GetThreadbyID(ID)
	// parsing the input json
	var NewThread = &models.Thread{}
	utils.ParseBody(r, NewThread)
	//populate new data
	db.Model(&ThreadDetails).Update("Title", NewThread.Title)
	db.Model(&ThreadDetails).Update("Body", NewThread.Body)
	db.Model(&ThreadDetails).Update("Tag", NewThread.Tag)
	db.Model(&ThreadDetails).Update("Likes", NewThread.Likes)
	db.Model(&ThreadDetails).Update("Dislikes", NewThread.Dislikes)

	res, _ := json.Marshal(ThreadDetails)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CountResponse(w http.ResponseWriter, r *http.Request) {
	// getting ID as param from /thread/comment/{id_thread}
	vars := mux.Vars(r)
	threadid := vars["id_thread"]
	ID, err := strconv.ParseInt(threadid, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	// call on models to accses database
	thread, db := models.GetThreadbyID(ID)
	count := models.ResponseCounter(ID)
	// update new count
	db.Model(&thread).Update("ResponseCount", count)

	res, _ := json.Marshal(thread)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
