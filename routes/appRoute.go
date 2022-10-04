package routes

import (
	"try_xml/controller"

	"github.com/gofiber/fiber/v2"
)

func AppRoutes(app *fiber.App) {

	accountEndpoint := app.Group("/account")
	accountEndpoint.Get("/get_user", controller.GetAccount)
}
