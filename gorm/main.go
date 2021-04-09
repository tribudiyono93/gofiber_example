package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/tribudiyono93/gofiber_example/gorm/book"
	"github.com/tribudiyono93/gofiber_example/gorm/database"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	initDatabase()
	sqlDB, err := database.DBConn.DB()
	if err != nil {
		panic("Failed get db")
	}
	defer sqlDB.Close()

	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Post("/api/v1/book", book.NewBook)
}

func initDatabase() {
	dbUser := "gouser"
	dbPass := "b7e6692c-79ec-11eb-9439-0242ac130002"
	dbHost := "127.0.0.1"
	dbName := "golang_rest_api_gin_gorm_mysql_jwt"

	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	database.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	database.DBConn.AutoMigrate(&book.BookAuthor{})

	sqlDb, err := database.DBConn.DB()
	if err != nil {
		panic("Failed to get db")
	}
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(10)
}
