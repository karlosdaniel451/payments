package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/karlosdaniel451/go-rest-api-template/api/controller"
)

func Setup(app *fiber.App, taskController *controller.UserController) {
	setupSwagger(app, "/docs/*")
	setupUserRouter(app, taskController)
}

func setupSwagger(app *fiber.App, path string) {
	app.Get("/docs/*", swagger.HandlerDefault)
}
