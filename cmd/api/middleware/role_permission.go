package middleware

import (
	"github.com/DanielChachagua/GestionCar/pkg/dependencies"
	"github.com/DanielChachagua/GestionCar/pkg/key"
	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func RolePermissionMiddleware(permissionName string) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user, ok := c.Locals("user").(*models.User)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(models.Response{
				Status:  false,
				Body:    nil,
				Message: "No autenticado",
			})
		}

		ctx := c.UserContext()
		depsTenant, ok := ctx.Value(key.TenantDBKey).(*dependencies.TenantApplication)
		if !ok {
			return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
				Status:  false,
				Body:    nil,
				Message: "Dependencias del tenant no disponibles",
			})
		}

		member, err := depsTenant.MemberController.MemberService.MemeberGetPermissionByUserID(user.ID)
		if err != nil {
			if errResp, ok := err.(*models.ErrorStruc); ok {
				return c.Status(errResp.StatusCode).JSON(models.Response{
					Status:  false,
					Body:    nil,
					Message: errResp.Message,
				})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(models.Response{
				Status:  false,
				Body:    nil,
				Message: "Error interno",
			})
		}

		if member.Role.Name == "admin" {
			return c.Next()
		}

		for _, permission := range member.Role.Permissions {
			if permission.Name == permissionName {
				return c.Next()
			}
		} 

		return c.Status(fiber.StatusForbidden).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "No autorizado",
		})
	}
}
