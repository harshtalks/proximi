package handlers

import (
	"app/utils"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// @Tags Auth
//
// @Router /auth/signup [post]
// @Summary For new User
// @Description Signup with username and password, you will receive a token which u will have to provide in the header of subsequent requests at /api gateway
// @Accept json
// @Produce json
// @failure 400,403,404,500 {object} ErrorModel
// @Param data body SigninInput true "username and password input"
func Signup(context *fiber.Ctx) error {
	requestBody := new(SigninInput)

	// checking if there is a parsing error
	if requestBodyError := context.BodyParser(requestBody); requestBodyError != nil {
		return context.Status(fiber.StatusBadRequest).JSON(ErrorModel{
			Status:     "error",
			StatusCode: fiber.StatusBadRequest,
			Message:    requestBodyError.Error(),
		})
	}

	var (
		password = requestBody.Password
		username = requestBody.Username
	)

	// Check if the password is strong

	if passwordStrengthErr := utils.IsPasswordStrong(password); passwordStrengthErr != nil {
		return context.Status(fiber.StatusForbidden).JSON(ErrorModel{
			Status:     "error",
			StatusCode: fiber.StatusForbidden,
			Message:    passwordStrengthErr.Error(),
		})
	}

	// we can proceed further

	user, userMutationError := utils.CreateUser(username, password)

	if userMutationError != nil {
		return context.Status(fiber.StatusNotFound).JSON(ErrorModel{
			Status:     "error",
			StatusCode: fiber.StatusNotFound,
			Message:    userMutationError.Error(),
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
		return context.Status(fiber.StatusInternalServerError).JSON(ErrorModel{
			Status:     "error",
			Message:    "Internal server error",
			StatusCode: fiber.StatusInternalServerError,
		})
	}

	return context.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":      "success",
		"status_code": fiber.StatusCreated,
		"user":        user,
		"token":       token,
	})

}
