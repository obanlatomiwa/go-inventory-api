package middlewares

import (
	"github.com/gofiber/fiber/v2"
	jwtMiddleWare "github.com/gofiber/jwt/v3"
	"github.com/obanlatomiwa/go-inventory-api/utils"
)

func CreateMiddleware() func(*fiber.Ctx) error {
	// create jwt middleware
	config := jwtMiddleWare.Config{
		SigningKey:   []byte(utils.GetValueFromConfigFile("JWT_SECRET_KEY")),
		ContextKey:   "jwt",
		ErrorHandler: jwtError,
	}

	return jwtMiddleWare.New(config)
}

func jwtError(c *fiber.Ctx, err error) error {
	// if the error is caused by a malformed jwt token return an error
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// if the error was caused by another error, return an error
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"message": err.Error(),
	})
}
