package main

import (
	"app/config"
	"app/database"
	"app/router"
	"fmt"

	_ "app/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

// @title Proximi
// @version 1.0
// @description Proximi service
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email harshpareek91@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:3000
// @BasePath /
func main() {

	// 1. Loading ENVs
	config.ConfigureENV()

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

// func getDistance() {
// 	apiKey := os.Getenv("BING_KEY")
// 	origin := "47.6062,-122.3321"      // Latitude and longitude of the origin
// 	destination := "37.7749,-122.4194" //
// 	url := fmt.Sprintf("https://dev.virtualearth.net/REST/v1/Routes/DistanceMatrix?origins=%s&destinations=%s&travelMode=driving&key=%s", origin, destination, apiKey)

// 	response, err := http.Get(url)

// 	if err != nil {
// 		fmt.Println("Error making the HTTP request:", err)
// 		return
// 	}

// 	body, err := ioutil.ReadAll(response.Body)
// 	if err != nil {
// 		fmt.Println("Error reading the response body:", err)
// 		return
// 	}

// 	fmt.Println(string(body))

// }
