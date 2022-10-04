package routes

import (
	"try_xml/controller"

	"github.com/gofiber/fiber/v2"

	//----- SWAGGER -----
	swagger "github.com/arsmn/fiber-swagger/v2"

	_ "try_xml/docs"
)

func AppRoutes(app *fiber.App) {

	//Swagger
	app.Get("/docs/*", swagger.HandlerDefault)

	accountEndpoint := app.Group("/account")
	accountEndpoint.Get("/get_user", controller.GetAccount)
}
