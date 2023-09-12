package handlers

import (
	"app/utils"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type SigninSuccessResponse struct {
	Status     string      `json:"status" example:"success"`
	StatusCode int         `json:"status_code" example:"200"`
	Data       SuccessData `json:"data"`
}

type SuccessData struct {
	Username  string    `json:"username"`
	UserID    uint      `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	Token     string    `json:"token"`
}

// @Tags Auth
//
// @Router /auth/signin [post]
// @Summary Basic token based Auth
// @Description Signin with username and password, you will receive a token which u will have to provide in the header of subsequent requests at /api gateway
// @Accept json
// @Produce json
// @failure 500,404,401,400 {object} ErrorModel
// @Success 200 {object} SigninSuccessResponse
// @Param data body SigninInput true "username and password input"
func Signin(context *fiber.Ctx) error {

	// Login handler

	signinInput := new(SigninInput)

	// parsing request body

	if signinInputError := context.BodyParser(signinInput); signinInputError != nil {
		return context.Status(fiber.StatusBadRequest).JSON(ErrorModel{
			Status:     "error",
			StatusCode: fiber.StatusBadRequest,
			Message:    signinInputError.Error(),
		})
	}

	var (
		username = signinInput.Username
	)

	// fetching user and error related to it

	user, userError := utils.GetUser(username)

	if userError == utils.NOTFOUND {
		return context.Status(fiber.StatusNotFound).JSON(ErrorModel{
			Status:     "error",
			Message:    fmt.Sprintf("given username %s is not found in our database.", username),
			StatusCode: fiber.StatusNotFound,
		})
	}

	if userError == utils.CUSTOMERROR {
		return context.Status(fiber.StatusInternalServerError).JSON(ErrorModel{
			Status:     "error",
			Message:    "Error in responding to the request",
			StatusCode: fiber.StatusInternalServerError,
		})
	}

	// check if the password is matched

	if isPasswordVerified := utils.VerifyPasswordHash(user.Password, signinInput.Password); !isPasswordVerified {
		return context.Status(fiber.StatusUnauthorized).JSON(ErrorModel{
			Status:     "error",
			Message:    fmt.Sprintf("password given for user %s does not match.", username),
			StatusCode: fiber.StatusUnauthorized,
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
		return context.Status(fiber.StatusInternalServerError).JSON(ErrorModel{
			Status:     "error",
			Message:    "Internal server error",
			StatusCode: fiber.StatusInternalServerError,
		})
	}

	// return response

	return context.Status(fiber.StatusOK).JSON(SigninSuccessResponse{
		Status:     "success",
		StatusCode: fiber.StatusOK,
		Data: SuccessData{
			Username:  user.Username,
			UserID:    user.ID,
			CreatedAt: user.CreatedAt,
			Token:     token,
		},
	})

}
