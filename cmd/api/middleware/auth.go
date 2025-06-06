package middleware

import (
	"github.com/DanielChachagua/GestionCar/pkg/database"
	"github.com/DanielChachagua/GestionCar/pkg/dependencies"
	"github.com/DanielChachagua/GestionCar/pkg/key"
	"github.com/DanielChachagua/GestionCar/pkg/models"
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

		ctx := c.UserContext()
		deps := ctx.Value(key.AppKey).(*dependencies.Application)

		claims, err := utils.VerifyToken(token)

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token inválido",
			})
		}

		userId := claims.(jwt.MapClaims)["id"].(string)
		user, err := deps.AuthController.AuthService.CurrentUser(userId)

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

		tenantID, ok := claims.(jwt.MapClaims)["tenant_id"].(string)
		if !ok {
			tenantID = ""
		}

		c.Locals("tenant", nil)

		if tenantID != "" {
			tenant, err := deps.TenantController.TenantService.TenantGetByID(userId, tenantID)
			if err != nil {
				return c.Status(fiber.StatusUnauthorized).JSON(models.Response{
					Status:  false,
					Body:    nil,
					Message: err.Error(),
				})
			}

			connection, err := utils.Decrypt(tenant.Connection)
			if err != nil {
				return c.Status(500).JSON(models.Response{
					Status:  false,
					Body:    nil,
					Message: err.Error(),
				})
			}

			db, err := database.GetTenantDB(connection)
			if err != nil {
				return c.Status(500).JSON(models.Response{
					Status:  false,
					Body:    nil,
					Message: err.Error(),
				})
			}

			mainDB := database.GetMainDB()

			ctx := c.UserContext()
			depsTenant := ctx.Value(key.TenantDBKey).(*dependencies.TenantApplication)
			depsTenant.SetDBTenantRepository(db, mainDB)
			c.Locals("tenant", tenant)
		}

		c.Locals("user", user)

		return c.Next()
	}
}
