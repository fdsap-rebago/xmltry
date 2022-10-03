package main

import (
	"try_xml/database"
	"try_xml/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Connect DB
	database.ConnectDatabase()
	// ----------

	app := fiber.New()
	routes.AppRoutes(app)

	app.Listen(":5555")

}
