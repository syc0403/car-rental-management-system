package utils

import "gorm.io/gorm"

// Paginate 分页封装
func Paginate(current int, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if current == 0 {
			current = 1
		}
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}
		offset := (current - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
