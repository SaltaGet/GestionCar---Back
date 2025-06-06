package middleware

import (
	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func TenantMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		tenantRaw := c.Locals("tenant")
		tenant, ok := tenantRaw.(*models.Tenant)
		if !ok || tenant == nil {
			return c.Status(401).JSON(models.Response{
				Status:  false,
				Body:    nil,
				Message: "login tenant is required",
			})
		}

		return c.Next()
	}
}
