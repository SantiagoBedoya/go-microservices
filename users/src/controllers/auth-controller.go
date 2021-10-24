package controllers

import (
	"fmt"
	"time"

	"github.com/SantiagoBedoya/go-microservices/users/src/database"
	"github.com/SantiagoBedoya/go-microservices/users/src/helpers"
	"github.com/SantiagoBedoya/go-microservices/users/src/models"
	"github.com/gofiber/fiber/v2"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confirm"] {
		return c.Status(400).JSON(fiber.Map{
			"message": "password do not match",
		})
	}
	user := models.User{
		FirstName: data["firstname"],
		LastName:  data["lastname"],
		Email:     data["email"],
	}

	user.SetPassword(data["password"])
	database.DB.Create(&user)

	helpers.PublishToEmailQueque(helpers.Data{
		Email:   user.Email,
		Subject: "Welcome to go-microservices",
		Message: fmt.Sprintf("Hi %s, we are so happy, welcome to the team", user.FullName()),
	})

	return c.Status(fiber.StatusCreated).JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user models.User

	database.DB.Where("email = ?", data["email"]).First(&user)

	if user.Id == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid credentials",
		})
	}
	if err := user.ComparePassword(data["password"]); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid credentials",
		})
	}
	token, err := helpers.GenerateJWT(user.Id)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "invalid credentials",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
}

func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
}
