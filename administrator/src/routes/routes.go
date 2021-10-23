package routes

import (
	"github.com/SantiagoBedoya/go-microservices/administrator/src/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("ok")
	})

	api := app.Group("api/administrator")
	api.Post("register")
	api.Post("login")

	adminAuthenticated := api.Use(middlewares.IsAuthenticated)
	adminAuthenticated.Get("user")
}
