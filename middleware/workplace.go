package middleware

import (
	"github.com/DanielChachagua/GestionCar/models"
	"github.com/DanielChachagua/GestionCar/services"
	"github.com/DanielChachagua/GestionCar/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func WorkplaceMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Get("X-Workplace-Token")

		if token == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token no proporcionado",
			})
		}

		claims, err := utils.VerifyWorkplaceToken(token)

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token inválido",
			})
		}

		workplaceId := claims.(jwt.MapClaims)["id"].(string)

		workplace, err := services.CurrentWorkplace(workplaceId)

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
		
		c.Locals("workplace", workplace)

		return c.Next()
	}
}