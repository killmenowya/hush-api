package models

import (
	"github.com/jinzhu/gorm"
	"github.com/killmenowya/hush-api/pkg/config"
)

var db *gorm.DB

type Comment struct {
	gorm.Model
	idComment int
	idThread  int
	idUser    int
	datetime  string
	Message   string
	Likes     int
	Dislikes  int
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Comment{})
}

func (c *Comment) CreateComment() *Comment {
	db.NewRecord(c)
	db.Create(&c)
	return c
}
