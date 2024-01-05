package models

import (
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/killmenowya/hush-api/pkg/utils"
)

type Saved struct {
	gorm.Model
	IdThread int
	IdUser   int
}

func (s *Saved) SaveThread() *Saved {
	// create new row to ensure no overlapping data
	db.NewRecord(s)
	// populate row with data
	db.Create(&s)
	return s
}

func DeleteSaved(idsaved int64) Saved {
	var s Saved
	// search by string for matching ID
	// delete row with matching ID
	db.Where("ID = ?", idsaved).Delete(&s)
	return s
}

func GetSaved(r *http.Request) []Saved {
	var saved []Saved
	// paginate and find data
	db.Scopes(utils.Paginate(r)).Find(&saved)
	return saved
}
