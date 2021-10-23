package routes

import (
	"github.com/SantiagoBedoya/go-microservices/users/src/controllers"
	"github.com/SantiagoBedoya/go-microservices/users/src/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("ok")
	})

	api := app.Group("api")

	api.Post("register", controllers.Register)
	api.Post("login", controllers.Login)

	authenticated := api.Use(middlewares.IsAutenticated)
	authenticated.Get("users", controllers.GetUsers)
	authenticated.Get("users/:id", controllers.GetUser)
	authenticated.Post("logout", controllers.Logout)
}
