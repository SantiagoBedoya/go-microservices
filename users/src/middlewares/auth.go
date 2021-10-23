package middlewares

import (
	"os"
	"strconv"

	"github.com/SantiagoBedoya/go-microservices/users/src/database"
	"github.com/SantiagoBedoya/go-microservices/users/src/models"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func IsAutenticated(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil || !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"messsage": "unauthorized",
		})
	}

	payload := token.Claims.(*jwt.StandardClaims)
	id, _ := strconv.Atoi(payload.Subject)
	var user models.User

	database.DB.Where("id = ?", id).First(&user)

	if user.Id == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "unauthorized",
		})
	}

	c.Context().SetUserValue("user", user)
	return c.Next()
}
