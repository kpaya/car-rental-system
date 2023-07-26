package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kpaya/car-rental-system/src/service"
)

func UserValidationMiddleware(c *fiber.Ctx) error {
	auth := c.GetReqHeaders()["Authorization"]
	if auth == "" {
		code := fiber.StatusBadRequest
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"msg":  "you must provide the authentication code",
			"code": code,
		})
	}

	if err := service.ValidateJWTToAccess(auth); err != nil {
		code := fiber.StatusBadRequest
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"msg":  err.Error(),
			"code": code,
		})
	}

	return c.Next()
}
