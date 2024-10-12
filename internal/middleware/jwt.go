package middleware

import (
	"errors"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)


func CheckJWT(c *fiber.Ctx) error{

	tokenStr := c.Get("Authorization")

	if tokenStr == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Missing or invalid Authorization header",
		})
	}

	tokenStr = tokenStr[len("Bearer "):]

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invaild algorithm for signing method")
		}

		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or expired token",
		})
	}

	if _, ok := token.Claims.(jwt.MapClaims); !ok || !token.Valid {

		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid or expired token",
		})

	}

	return c.Next()
}