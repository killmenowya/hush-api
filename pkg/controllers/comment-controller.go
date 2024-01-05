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

var NewComment models.Comment

// CREATE
func CreateComment(w http.ResponseWriter, r *http.Request) {
	// read input body and parse it for use
	CreateComment := &models.Comment{}
	utils.ParseBody(r, CreateComment)
	// call on models to access database
	c := CreateComment.CreateComment()
	res, _ := json.Marshal(c)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// DELETE
func DeleteComment(w http.ResponseWriter, r *http.Request) {
	// getting id as param from /comment/{id_comment}
	vars := mux.Vars(r)
	commentid := vars["id_comment"]
	ID, err := strconv.ParseInt(commentid, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	// call on models to access database
	comment := models.DeleteComment(ID)
	res, _ := json.Marshal(comment)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

// READ
func GetComment(w http.ResponseWriter, r *http.Request) {
	// getting id as param from /comment/{id_thread}
	vars := mux.Vars(r)
	threadid := vars["id_thread"]
	ID, err := strconv.ParseInt(threadid, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	// call on models to access database
	comment := models.GetCommentbyID(ID)
	res, _ := json.Marshal(comment)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

// UPDATE
func EditComment(w http.ResponseWriter, r *http.Request) {
	// getting the id as param from /commemt/{id_comment}
	vars := mux.Vars(r)
	commentid := vars["id_comment"]
	ID, err := strconv.ParseInt(commentid, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	// getting the original comment and db address
	CommentDetails, db := models.GetComment(ID)
	// parsing the input json
	var NewComment = &models.Comment{}
	utils.ParseBody(r, NewComment)
	// populate new data
	db.Model(&CommentDetails).Update("Message", NewComment.Message)
	db.Model(&CommentDetails).Update("Likes", NewComment.Likes)
	db.Model(&CommentDetails).Update("Dislikes", NewComment.Dislikes)

	res, _ := json.Marshal(CommentDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
