package models

import (
	"github.com/jinzhu/gorm"
)

type Tag struct {
	gorm.Model
	Tag     string
	Counter int
}

func (t *Tag) CreateTags() *Tag {
	// create new row to ensure no overlappig data
	db.NewRecord(t)
	// populate row with data
	db.Create(&t)
	return t

}

func GetTags() []Tag {
	var t []Tag
	// order the data desc
	// limit the return to 5
	// return top 5 tags with highest counter
	db.Order("counter DESC").Limit(5).Find(&t)
	return t
}

func GetTagbyName(TagName string) (*Tag, *gorm.DB) {
	var gettag Tag
	// search by string for matching ID
	// return row with matched ID
	db := db.Where("tag = ?", TagName).Find(&gettag)
	return &gettag, db
}

func Counter(tag string) int64 {
	var count int64
	// search [threads] table by string for matching tag
	// return count of matching row
	db.Table("threads").Where("tag = ?", tag).Count(&count)
	return count
}
