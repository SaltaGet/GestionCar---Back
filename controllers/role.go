package controllers

import (
	"github.com/DanielChachagua/GestionCar/models"
	"github.com/DanielChachagua/GestionCar/services"
	"github.com/gofiber/fiber/v2"
)

// GetRolesWorkplace godoc
//	@Summary		Retrieve roles for a user in a specific workplace
//	@Description	This function fetches roles based on the user's role and workplace identifier
//              from the context. It requires both user and workplace information to be present
//              in the request context.
//	@Tags			Role
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			X-Workplace-Token	header		string			true	"Workplace Token"
//	@Success		200					{object}	models.Response	"Roles retrieved successfully"
//	@Failure		400					{object}	models.Response	"Bad request if user or workplace is missing"
//	@Failure		500					{object}	models.Response	"Internal server error on failure"
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