package router

import (
	"app/handlers"
	"app/middleware"

	"github.com/gofiber/fiber/v2"
)

func Router(app *fiber.App) {
	// making routes

	// applying rate limiter

	app.Use(middleware.RateLimiter())

	// Auth group apis -> /auth/*
	authGroup := app.Group("/auth")

	authGroup.Post("/signup", handlers.Signup)
	authGroup.Post("/signin", handlers.Signin)

	// api groups

	apiGroup := app.Group("/api", middleware.ProtectRoutes())

	apiGroup.Get("/verify", handlers.Verify)

	// api for the businesses
	apiGroup.Get("/businesses", handlers.Businesses)
	apiGroup.Get("/businesses/:businessId", handlers.Business)
}
