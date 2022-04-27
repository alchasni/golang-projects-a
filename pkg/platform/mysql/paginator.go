package mysql

import "gorm.io/gorm"

type paginator struct {
	minPageSize int
	maxPageSize int
}

func (p paginator) paginate(db *gorm.DB, pageNo, pageSize int) *gorm.DB {
	if p.minPageSize > 0 && pageSize < p.minPageSize {
		pageSize = p.minPageSize
	} else if p.maxPageSize > 0 && pageSize > p.maxPageSize {
		pageSize = p.maxPageSize
	}
	db = db.Limit(pageSize)

	if pageNo > 0 {
		db = db.Offset((pageNo - 1) * pageSize)
	}

	return db
}
