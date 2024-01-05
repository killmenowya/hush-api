package models

import "github.com/jinzhu/gorm"

//"time"

type Comment struct {
	gorm.Model
	IdThread int
	IdUser   int
	Message  string
	Likes    int
	Dislikes int
}

func (c *Comment) CreateComment() *Comment {
	// create new row to ensure no overlapping data
	db.NewRecord(c)
	// populate row with data
	db.Create(&c)
	return c
}

func DeleteComment(idcomment int64) Comment {
	var c Comment
	// search by string for matching ID
	// delete row with matching ID
	db.Where("ID = ?", idcomment).Delete(c)
	return c
}

func GetCommentbyID(idthread int64) []Comment {
	var c []Comment
	// sort by likes desc
	// search by string for matching thread_ID
	// return rows into array to read
	db.Order("likes DESC").Where("id_thread = ?", idthread).Find(&c)
	return c
}

func GetComment(idcomment int64) (*Comment, *gorm.DB) {
	var comment Comment
	// search by string for matching ID
	// return row with matched ID
	db := db.Where("ID = ?", idcomment).Find(&comment)
	return &comment, db

}
