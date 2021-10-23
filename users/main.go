package main

import (
	"github.com/SantiagoBedoya/go-microservices/users/src/database"
	"github.com/SantiagoBedoya/go-microservices/users/src/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()
	database.AutoMigrate()

	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("ok")
	})
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.Setup(app)

	app.Listen(":8000")
}
