package handlers

import (
	"app/database"
	"app/models"
	"app/service"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

// @Tags Businesses
//
// @Router /api/businesses/:businessId [get]
// @Summary get details of a business
// @description.markdown individualBusiness.md
//
// @Param businessId path string true "Id of the business"
// @Param lat query string false "latitude of your current location"
// @Param long query string false "longitude of your current location"
// @Param travelMode query string false "your mode of travel" Enums("driving", "walking", "public transit")
//
// @failure 404,500 {object} ErrorModel
// @Success 200 {object} map[string]interface{}
//
// @Security ApiKeyAuth
func Business(context *fiber.Ctx) error {
	businessId := context.Params("businessId")

	// if user needs to see the distance between them and the business
	latitude := context.Query("lat")
	longitude := context.Query("long")
	travelMode := context.Query("travelMode")

	var business models.Businesses

	businessErr := database.Database.DB.Where("business_id = ?", businessId).First(&business).Error

	if businessErr != nil {
		return context.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":      "error",
			"message":     businessErr.Error(),
			"status_code": fiber.StatusNotFound,
		})
	}

	// that means we have business.

	var distanceObject *fiber.Map
	var distanceBool bool
	var message string

	if latitude != "" && longitude != "" {
		// set default travelMode to driving if no mode given or invalid mode given
		if travelMode == "" || !service.IsTravelModeValid(travelMode) {
			travelMode = "driving"
		}

		// origin codes
		originCoords := service.Coordinate{
			Lat:  latitude,
			Long: longitude,
		}

		// destination codes
		destinationCoords := service.Coordinate{
			Lat:  strconv.FormatFloat(business.Latitude, 'f', -1, 64),
			Long: strconv.FormatFloat(business.Longitude, 'f', -1, 64),
		}

		// error variable
		var err error

		// getting distance
		distanceObject, err = service.GetDistance(&originCoords, &destinationCoords, travelMode)
		message = "Successfully got the data"
		distanceBool = true
		if err != nil {
			distanceBool = false
			message = err.Error()
		}

	}

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      "success",
		"status_code": fiber.StatusOK,
		"result": fiber.Map{
			"business": business,
			"distance": fiber.Map{
				"success": distanceBool,
				"message": message,
				"travel":  distanceObject,
			},
		},
	})
}
