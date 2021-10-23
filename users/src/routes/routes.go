package routes

import (
	"github.com/SantiagoBedoya/go-microservices/users/src/controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("ok")
	})

	api := app.Group("api")

	api.Post("register", controllers.Register)
	api.Post("login", controllers.Login)
	// api.Post("users")
	// api.Post("users/:id")

	// authenticated := api.Use()
	// authenticated.Get("user/:scope")
	// authenticated.Post("logout")
	// authenticated.Post("users/info")
	// authenticated.Post("users/password")
}
