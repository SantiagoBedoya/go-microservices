package main

import (
	"github.com/SantiagoBedoya/go-microservices/administrator/src/helpers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	helpers.SetupAMQP()
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	app.Listen(":8000")
}
