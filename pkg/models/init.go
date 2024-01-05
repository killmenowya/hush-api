package models

import (
	"github.com/jinzhu/gorm"
	"github.com/killmenowya/hush-api/pkg/config"
)

var db *gorm.DB

func init() {
	// to create a connection with mysql database
	config.Connect()
	// get database address
	db = config.GetDB()
	// migrate schema to ensure schema up to date
	db.AutoMigrate(&Comment{})
	db.AutoMigrate(&Thread{})
	db.AutoMigrate(&Tag{})
	db.AutoMigrate(&Saved{})

}
