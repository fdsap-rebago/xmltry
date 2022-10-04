package main

import (
	"try_xml/middleware/database"
	"try_xml/routes"

	"github.com/gofiber/fiber/v2"
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

	app := fiber.New()
	routes.AppRoutes(app)

	app.Listen(":5555")

}
