package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/tribudiyono93/gofiber_example/fiber-rest-api/database"
	"github.com/tribudiyono93/gofiber_example/fiber-rest-api/router"
	"log"
	"os"
	"os/signal"
)

func main() {
	//load config first time
	if err := godotenv.Load(); err != nil {
		panic("Failed to load env file")
	}

	//connect and defer close database
	database.Connect()

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	//register router
	router.Routes(app)

	//graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		log.Println("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	if err := app.Listen(os.Getenv("SERVER_PORT")); err != nil {
		log.Panic(err)
	}

	//clean up resource
	log.Println("clean up resource")
	database.Close()
	log.Println("successfully graceful shutdown")
}
