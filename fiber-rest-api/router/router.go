package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tribudiyono93/gofiber_example/fiber-rest-api/handler"
	"net/http"
)

func Routes(app *fiber.App) {
	//sample
	app.Get("/", handler.Hello)
	app.Get("/error", handler.Error)

	//root api group
	api := app.Group("/api/v1")

	//auth
	auth := api.Group("/auth")
	auth.Post("/register", handler.Register)
	auth.Post("/login", handler.Login)


	// 404 Handler
	app.Use(func(c *fiber.Ctx) error { return c.SendStatus(http.StatusNotFound) })
}