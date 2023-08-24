package middleware

import (
	"os"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
)

func ProtectRoutes() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
		ErrorHandler: jwtErr,
	})
}

func jwtErr(context *fiber.Ctx, jwtError error) error {
	if jwtError.Error() == "Missing or malformed JWT" {
		return context.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":      "error",
			"error":       jwtError.Error(),
			"status_code": fiber.StatusBadRequest,
		})
	}

	return context.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"status":      "error",
		"message":     "you are not authorized to access this resource. please login.",
		"status_code": fiber.StatusUnauthorized,
	})
}
