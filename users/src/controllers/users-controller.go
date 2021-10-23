package controllers

import (
	"strconv"

	"github.com/SantiagoBedoya/go-microservices/users/src/database"
	"github.com/SantiagoBedoya/go-microservices/users/src/models"
	"github.com/gofiber/fiber/v2"
)

func GetUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "id param must be a number",
		})
	}
	user := models.User{
		Id: uint(id),
	}
	database.DB.Find(&user)

	if user.Email == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "user not found",
		})
	}

	return c.JSON(user)
}

func GetUsers(c *fiber.Ctx) error {
	var users []models.User
	database.DB.Find(&users)
	return c.Status(fiber.StatusOK).JSON(users)
}
