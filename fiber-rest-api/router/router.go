package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tribudiyono93/gofiber_example/fiber-rest-api/handler"
)

func Routes(app *fiber.App) {
	app.Get("/", handler.Hello)
}