package middleware

import (
	"simple-attendance/pkg/config"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
)

var ACCESS_TOKEN_KEY string = config.GetKeyConfig("ACCESS_TOKEN_KEY")

func ProtectedJWT() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:    []byte(ACCESS_TOKEN_KEY),
		SigningMethod: "HS256",
		AuthScheme:    "Bearer",
		ErrorHandler:  jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
		c.Set("content-type", "application/json; charset=utf-8")
		return nil
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}
