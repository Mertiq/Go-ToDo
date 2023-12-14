package main

import (
	"Go-ToDo/postgres"
	"Go-ToDo/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	postgres.Connect()

	app := fiber.New()
	app.Use(cors.New())

	routes.AddRoutes(app)
	app.Listen(":3000")

}
