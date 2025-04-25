package main

import (
	"app/database"
	"app/router"
	"log"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Suver",
	})
	// app.Use(cors.New())

	database.ConnectDB()

	// app.Get("/test", handler.Hello)

	router.SetupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
