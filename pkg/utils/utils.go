package utils

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
)

func ParseBody(r *http.Request, x interface{}) {
	// parse input body into readable data for api
	if body, err := io.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	// paginated database
	return func(db *gorm.DB) *gorm.DB {
		q := r.URL.Query()
		page, _ := strconv.Atoi(q.Get("page"))
		if page <= 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(q.Get("page_size"))
		switch {
		case pageSize > 10:
			pageSize = 10
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
