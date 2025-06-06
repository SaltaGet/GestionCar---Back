package middleware

import (
	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func AdminTenantMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tenant, ok := c.Locals("tenant").(*models.Tenant)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(models.Response{
				Status:  false,
				Body:    nil,
				Message: "Teanant No autenticado",
			})
		}
		
		if tenant.UserTenants[0].IsAdmin {
			return c.Next()
		}

		return c.Status(fiber.StatusForbidden).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "No autorizado",
		})
	}
}