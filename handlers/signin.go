package handlers

import (
	"app/utils"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type SigninInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Signin(context *fiber.Ctx) error {
	// Login handler

	signinInput := new(SigninInput)

	// parsing request body

	if signinInputError := context.BodyParser(signinInput); signinInputError != nil {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      "error",
			"status_code": fiber.StatusBadRequest,
			"message":     signinInputError.Error(),
		})
	}

	var (
		username = signinInput.Username
	)

	// fetching user and error related to it

	user, userError := utils.GetUser(username)

	if userError == utils.NOTFOUND {
		return context.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":      "error",
			"message":     fmt.Sprintf("given username %s is not found in our database.", username),
			"status_code": fiber.StatusNotFound,
		})
	}

	if userError == utils.CUSTOMERROR {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":      "error",
			"message":     "Error in responding to the request",
			"status_code": fiber.StatusInternalServerError,
		})
	}

	// check if the password is matched

	if isPasswordVerified := utils.VerifyPasswordHash(user.Password, signinInput.Password); !isPasswordVerified {
		return context.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":      "error",
			"message":     fmt.Sprintf("password given for user %s does not match.", username),
			"status_code": fiber.StatusUnauthorized,
		})
	}

	// JWT stuff

	// jwt claim

	claims := jwt.MapClaims{
		"username": username,
		"user_id":  user.ID,
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	jwtResponse := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, tokenError := jwtResponse.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if tokenError != nil {
		return context.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":      "error",
			"message":     "Internal server error",
			"status_code": fiber.StatusInternalServerError,
		})
	}

	// return response

	return context.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":      "success",
		"status_code": fiber.StatusOK,
		"data": fiber.Map{
			"username":   user.Username,
			"user_id":    user.ID,
			"created_at": user.CreatedAt,
			"token":      token,
		},
	})

}
