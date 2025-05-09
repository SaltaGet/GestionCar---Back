package middleware

import (
	"github.com/DanielChachagua/GestionCar/models"
	"github.com/DanielChachagua/GestionCar/utils"
	"github.com/gofiber/fiber/v2"
)

func RoleAuthMiddleware(roles []string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user, ok := c.Locals("user").(*models.User)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(models.Response{
				Status:  false,
				Body:    nil,
				Message: "Unauthorized",
			})
		}
		
		if utils.Contains(roles, user.Role) {
			return c.Next()
		}

		return c.Status(fiber.StatusUnauthorized).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Unauthorized",
		})
	}
}