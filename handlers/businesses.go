package handlers

import (
	"app/database/scopes"
	"app/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/mmcloughlin/geohash"
)

func Businesses(context *fiber.Ctx) error {
	// Get the Latitude and Longitude from the queries

	latitude := context.Query("lat")

	if latitude == "" {
		return context.Status(fiber.StatusPreconditionFailed).JSON(fiber.Map{
			"status":      "error",
			"message":     "Latitude is missing from the query",
			"status_code": fiber.StatusPreconditionFailed,
		})
	}

	longitude := context.Query("long")

	if longitude == "" {
		return context.Status(fiber.StatusPreconditionFailed).JSON(fiber.Map{
			"status":      "error",
			"message":     "Longitude is missing from the query",
			"status_code": fiber.StatusPreconditionFailed,
		})
	}

	// range described

	radiusQuery := context.Query("range")

	radius, radiusErr := strconv.ParseFloat(radiusQuery, 64)

	if radiusErr != nil {
		radius = 2
	}

	hashLength := utils.GetGeoHashLength(radius)

	// geo hash

	parsedLatitude, parsedLatErr := strconv.ParseFloat(latitude, 64)
	parsedLongitude, parsedLongErr := strconv.ParseFloat(longitude, 64)

	if parsedLatErr != nil || parsedLongErr != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      "error",
			"message":     "Error Parsing the Latitude and Longitude Values",
			"status_code": fiber.StatusBadRequest,
		})
	}

	// Generating location Hash
	locationHash := geohash.EncodeWithPrecision(parsedLatitude, parsedLongitude, uint(hashLength))

	// Businesses
	businesses, businessCounts, queryError := utils.GetBusinesses(locationHash, context)

	if queryError == utils.CUSTOMERROR {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":      "error",
			"message":     "Error in responding to the request",
			"status_code": fiber.StatusInternalServerError,
		})
	}

	// for the avaiable rows, we will see if next page is available.

	page, perPage := scopes.PaginationParams(context)

	utils.IsNextPageAvailable(page, int(businessCounts.Count), perPage, context)

	// this means that the api gives us the correct result

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      "success",
		"status_code": fiber.StatusOK,
		"lat":         latitude,
		"long":        longitude,
		"result": fiber.Map{
			"businesses": businesses,
			"length":     len(businesses),
		},
	})
}
