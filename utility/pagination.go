package utility

import (
	"gorm.io/gorm"
	"math"
)

func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func Roundf(x float64) float64 {
	t := math.Trunc(x)
	if math.Abs(x-t) >= 0.0001 {
		return t + math.Copysign(1, x)
	}
	return t
}