package book

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/tribudiyono93/gofiber_example/gorm/database"
	"gorm.io/gorm"
)

type BookAuthor struct {
	gorm.Model
	Title  string `json:"name"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) error {
	fmt.Printf("db memory location %v\n",&database.DBConn)
	var book []BookAuthor
	database.DBConn.Find(&book)
	return c.JSON(book)
}

func NewBook(c *fiber.Ctx) error {
	db := database.DBConn
	book := new(BookAuthor)
	if err := c.BodyParser(book); err != nil {
		return c.Status(503).SendString(err.Error())
	}
	db.Create(&book)
	return c.JSON(book)
}
