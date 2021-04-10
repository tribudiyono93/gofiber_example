package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tribudiyono93/gofiber_example/fiber-rest-api/handler"
	"github.com/tribudiyono93/gofiber_example/fiber-rest-api/middleware"
	"net/http"
)

var allModuleAllRoles = map[string][]string {
	"moduleA": {"ADMIN","STAFF"},
	"moduleB": {"ADMIN","STAFF"},
}

var moduleCAdmin = map[string][]string {
	"moduleC": {"ADMIN"},
}

func Register(app *fiber.App) {
	//sample
	app.Get("/", handler.Hello)
	app.Get("/error", handler.Error)

	//root api group
	api := app.Group("/api/v1")

	//auth
	auth := api.Group("/auth")
	auth.Post("/register", handler.Register)
	auth.Post("/login", handler.Login)

	secure := api.Group("/secure", middleware.ValidateJWT())
	secure.Get("/test", middleware.HasModuleAndRole(moduleCAdmin), handler.Hello)
	secure.Get("/test2", middleware.HasModuleAndRole(allModuleAllRoles), handler.Hello)

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error { return c.SendStatus(http.StatusNotFound) })
}