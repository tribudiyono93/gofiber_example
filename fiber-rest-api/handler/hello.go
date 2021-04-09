package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tribudiyono93/gofiber_example/fiber-rest-api/entity"
	"github.com/tribudiyono93/gofiber_example/fiber-rest-api/response"
	"net/http"
	"time"
)

func Hello(c *fiber.Ctx) error {
	var book = entity.Book{
		ID: 1,
		Title: "Book Title",
		Author: "Author",
		Rating: 10,
		Base: entity.Base{
			CreatedBy: "Tri",
			CreatedAt: time.Now(),
			UpdatedBy: "Tri",
			UpdatedAt: time.Now(),
		},
	}
	return c.Status(http.StatusOK).JSON(book)
}

func Error(c *fiber.Ctx) error {
	return c.Status(http.StatusBadRequest).JSON(response.Error{
		Code: entity.UserNotFound,
		Message: entity.StatusText[entity.UserNotFound],
	})
}