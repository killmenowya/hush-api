package controllers

import (
	"encoding/json"
	//"fmt"
	//"github.com/gorilla/mux"
	"net/http"
	//"strconv"
	"github.com/killmenowya/hush-api/pkg/models"
	"github.com/killmenowya/hush-api/pkg/utils"
)

var NewComment models.Comment

func CreateComment(w http.ResponseWriter, r *http.Request) {
	CreateComment := &models.Comment{}
	utils.ParseBody(r, CreateComment)
	b := CreateComment.CreateComment()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
