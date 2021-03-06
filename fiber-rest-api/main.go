package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/joho/godotenv"
	"github.com/tribudiyono93/gofiber_example/fiber-rest-api/connection"
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

	//connect and defer close connection
	connection.ConnectDB()

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Format: "${pid} ${locals:requestid} ${status} - ${method} ${path}\n",
	}))
	app.Use(recover.New())

	//register router
	router.Register(app)

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
	connection.CloseDB()
	log.Println("successfully graceful shutdown")
}
