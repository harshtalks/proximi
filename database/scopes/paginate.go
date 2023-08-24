package scopes

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func PaginationParams(context *fiber.Ctx) (int, int) {
	pageQuery := context.Query("page")

	page, pageErr := strconv.Atoi(pageQuery)

	if pageErr != nil || page <= 0 {
		page = 1
	}

	// Get business counts per page query
	perPageQuery := context.Query("perPage")

	perPage, perPageErr := strconv.Atoi(perPageQuery)

	switch {
	case perPageErr != nil, perPage <= 10:
		perPage = 10
	case perPage > 100:
		perPage = 100
	}

	return page, perPage
}

// Paginate Scope
func Paginate(context *fiber.Ctx) func(db *gorm.DB) *gorm.DB {

	return func(db *gorm.DB) *gorm.DB {

		page, perPage := PaginationParams(context)

		// offset
		offset := (page - 1) * perPage

		// returning paginated DB
		return db.Offset(offset).Limit(perPage)
	}

}
