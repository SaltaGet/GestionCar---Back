package controllers

import (
	"github.com/DanielChachagua/GestionCar/models"
	"github.com/DanielChachagua/GestionCar/services"
	"github.com/gofiber/fiber/v2"
)

func GetRolesWorkplace( c*fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	if user == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "User is required",
		})
	}

	workplace := c.Locals("workplace").(*models.Workplace)
	if workplace == nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status":  false,
			"message": "Workplace is required",
		})
	}

	roles, err := services.GetRoleAll(user.Role, workplace.Identifier)
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
	
	return c.JSON(fiber.Map{
		"status":  true,
		"message": "Roles retrieved successfully",
		"data":    roles,
	})
}