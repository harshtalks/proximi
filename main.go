package main

import (
	"app/database"
	"app/router"
	"fmt"

	_ "app/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

// @title		Proximi
// @description.markdown api
// @version	1.0
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.email	harshpareek91@gmail.com
// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath		/
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	// 1. Loading ENVs
	// config.ConfigureENV()

	// 2. Creating an instance of fiber app
	fmt.Println("New Fiber app is instanitiated...")
	app := fiber.New()

	// 3. Adding middlewares
	fmt.Println("Middleware logger registered...")
	app.Use(logger.New())

	// 4. Opening Database instance, connecting to it

	database.ConnectDatabase()

	// 5. Migration of database models

	// database.Migrate()

	// database.SeedDatabase()
	// 6. Routes

	router.Router(app)
	// swagger
	app.Get("/docs/*", swagger.HandlerDefault) // default

	// listening to the server
	app.Listen(":3000")

}
