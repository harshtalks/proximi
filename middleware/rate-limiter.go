package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

func RateLimiter() fiber.Handler {
	return limiter.New(limiter.Config{
		Max:        30,
		Expiration: 30 * time.Second,
		LimitReached: func(context *fiber.Ctx) error {
			return context.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"status":      "error",
				"message":     "too many requests",
				"status_code": fiber.StatusTooManyRequests,
			})
		},
	})
}
