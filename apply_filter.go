package helpers

import (
	"fmt"
	"strings"
	"gorm.io/gorm"
)

type Search struct {
	Keyword string `query:"keyword"`
	Column  string `query:"column"`
}
type Pagination struct {
	Page    int   `query:"page"`
	PerPage int   `query:"per_page"`
	Total   int64 `query:"total"`
}

func ApplySearchAndPagination(db *gorm.DB, filter Search, pagination *Pagination, model interface{}) (*gorm.DB, error) {
	db = ApplySearch(db, filter)
	db = ApplyPagination(db, pagination, model)
	return db, nil
}
func ApplySearch(db *gorm.DB, filter Search) *gorm.DB {
	if filter.Keyword == "" || filter.Column == "" {
		return db
	}
	columns := strings.Split(filter.Column, ",")
	var query string

	if len(columns) > 0 {
		var args []interface{}

		for _, column := range columns {
			if query != "" {
				query += " OR "
			}
			query += fmt.Sprintf("%s ILIKE ?", column)
			args = append(args, "%"+filter.Keyword+"%")
		}

		return db.Where(query, args...)
	} else {
		return db.Where(fmt.Sprintf("%s ILIKE ?", filter.Column), "%"+filter.Keyword+"%")
	}
}
func ApplyPagination(db *gorm.DB, pagination *Pagination, model interface{}) *gorm.DB {
	var total int64
	err := db.Model(model).Count(&total).Error
	if err != nil {
		return nil
	}

	pagination.Total = total
	if pagination.Page < 1 {
		pagination.Page = 1
	}
	if pagination.PerPage < 1 {
		pagination.PerPage = 10
	}

	return db.Offset((pagination.Page - 1) * pagination.PerPage).Limit(pagination.PerPage)
}
