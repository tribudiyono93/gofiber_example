package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tribudiyono93/gofiber_example/fiber-rest-api/handler"
	"net/http"
)

func Routes(app *fiber.App) {
	app.Get("/", handler.Hello)
	app.Get("/error", handler.Error)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusNotFound)
	})


}