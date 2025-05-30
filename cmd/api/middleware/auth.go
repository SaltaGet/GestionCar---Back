package middleware

import (
	"github.com/DanielChachagua/GestionCar/pkg/database"
	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/DanielChachagua/GestionCar/pkg/services"
	"github.com/DanielChachagua/GestionCar/pkg/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("Authorization")
		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token no proporcionado",
			})
		}

		tenantDB, err := database.GetTenantDB("default")
		if err != nil {
			return c.Status(500).JSON(models.Response{
				Status:  false,
				Message: "Error de conexión al tenant",
			})
		}

		ctx := c.UserContext()
		deps := ctx.Value(key.AppKey).(*dependencies.Application)
		deps.SetDBRepository(tenantDB)

		claims, err := utils.VerifyToken(token)

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token inválido",
			})
		}

		userId := claims.(jwt.MapClaims)["id"].(string)

		user, err := services.CurrentUser(userId)

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
		
		c.Locals("user", user)

		return c.Next()
	}
}