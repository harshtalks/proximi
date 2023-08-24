package handlers

import (
	"app/utils"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Signup(context *fiber.Ctx) error {
	requestBody := new(SigninInput)

	// checking if there is a parsing error
	if requestBodyError := context.BodyParser(requestBody); requestBodyError != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      "error",
			"status_code": fiber.StatusBadRequest,
			"message":     requestBodyError.Error(),
		})
	}

	var (
		password = requestBody.Password
		username = requestBody.Username
	)

	// Check if the password is strong

	if passwordStrengthErr := utils.IsPasswordStrong(password); passwordStrengthErr != nil {
		return context.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"status":      "error",
			"status_code": fiber.StatusForbidden,
			"message":     passwordStrengthErr.Error(),
		})
	}

	// we can proceed further

	user, userMutationError := utils.CreateUser(username, password)

	if userMutationError != nil {
		return context.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":      "error",
			"status_code": fiber.StatusNotFound,
			"message":     userMutationError.Error(),
		})
	}

	// user successfully created.

	// JWT stuff

	claims := jwt.MapClaims{
		"username": username,
		"user_id":  user.ID,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	jwtResponse := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, tokenErr := jwtResponse.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if tokenErr != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":      "error",
			"message":     "Internal server error",
			"status_code": fiber.StatusInternalServerError,
		})
	}

	return context.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":      "success",
		"status_code": fiber.StatusCreated,
		"user":        user,
		"token":       token,
	})

}
