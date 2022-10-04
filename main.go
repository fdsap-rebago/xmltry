package main

import (
	"try_xml/middleware/database"
	"try_xml/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// @title API - XML
// @version 1.0
// @description FDSAP swagger template
// @termsOfService http://swagger.io/terms/
// @contact.name FDSAP Support
// @contact.email raymond.bago@fortress-asya.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:5555
// @BasePath /

func main() {
	// Connect DB
	database.ConnectDatabase()
	// ----------

	// Declare & initialize fiber
	app := fiber.New(fiber.Config{
		UnescapePath: true,
	})

	// Configure application CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	routes.AppRoutes(app)

	app.Listen(":5555")

}
