package utils

import (
	"app/database"
	"app/database/scopes"
	"app/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type RowCount struct {
	Count int64
}

func GetBusinesses(geohash string, context *fiber.Ctx) ([]models.Businesses, *RowCount, ErrorUser) {
	// Database Instance
	DB := database.Database.DB

	// Variables
	var businesses []models.Businesses
	businessCounts := &RowCount{}
	var rowCounts int64

	// Main Query
	if queryErr := DB.Scopes(scopes.Paginate(context)).Where("geo_hash LIKE ?", geohash+"%s").Find(&businesses).Select("SQL_CALC_FOUND_ROWS *").Error; queryErr != nil {
		return nil, nil, CUSTOMERROR
	}

	// Query to count the total values
	if countErr := DB.Model(&models.Businesses{}).Where("geo_hash LIKE ?", geohash+"%s").Count(&rowCounts).Error; countErr != nil {
		return nil, nil, CUSTOMERROR
	}

	// Assigning
	businessCounts.Count = rowCounts

	return businesses, businessCounts, NOERROR
}

// This function sets the header that can be accessed by clients
// better UI
func IsNextPageAvailable(currentPage int, totalCount int, perPage int, context *fiber.Ctx) {
	totalPages := int(totalCount) / perPage

	// Determine if next page exists
	nextPageExists := currentPage < totalPages

	// Set custom header indicating if next page is available
	context.Set("X-Next-Page-Available", strconv.FormatBool(nextPageExists))
}
