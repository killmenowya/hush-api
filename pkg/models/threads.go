package models

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/killmenowya/hush-api/pkg/utils"
)

//"time"

type Thread struct {
	gorm.Model
	IdUser        int
	Title         string
	Body          string
	Tag           string
	Likes         int
	Dislikes      int
	ResponseCount int
	IsHidden      bool
}

func (t *Thread) CreateThread() *Thread {
	// create new row to ensure no overlapping data
	db.NewRecord(t)
	// populate row with data
	db.Create(&t)
	return t
}

func DeleteThread(idthread int64) Thread {
	var t Thread
	// search by string for matching ID
	// delete row with matching ID
	db.Where("ID = ?", idthread).Delete(t)
	return t
}

func GetThread(r *http.Request) []Thread {
	var t []Thread
	// paginate and find data
	db.Scopes(utils.Paginate(r)).Find(&t)
	return t
}

func GetThreadbyUser(r *http.Request, iduser int64) []Thread {
	var t []Thread
	// get the output DB as paginated
	PaginatedDB := utils.Paginate(r)(db.Model(&Thread{}))
	// search by string for matching ID
	// return row with matched ID
	PaginatedDB.Where("id_user = ?", iduser).Find(&t)
	return t
}

func GetThreadbyID(idthread int64) (*Thread, *gorm.DB) {
	var getthread Thread
	// search by string for matching ID
	// return row with matched ID
	db := db.Where("ID = ?", idthread).Find(&getthread)
	return &getthread, db
}

func GetThreadbyTag(r *http.Request, Tag string) []Thread {
	var getthread []Thread
	// get the output as paginated
	PaginatedDB := utils.Paginate(r)(db.Model(&Thread{}))
	// search by string for matching ID
	// return row with matched ID
	PaginatedDB.Where("tag = ?", Tag).Find(&getthread)
	return getthread
}

func ResponseCounter(idthread int64) int64 {
	var count int64
	db.Table("comments").Where("id_thread = ?", idthread).Count(&count)
	return count
}
