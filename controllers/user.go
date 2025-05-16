package controllers

import (
	"github.com/DanielChachagua/GestionCar/models"
	"github.com/DanielChachagua/GestionCar/services"
	"github.com/gofiber/fiber/v2"
)

// CreateUser godoc
//	@Summary		Create User
//	@Description	Creates a new user.
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			userCreate	body		models.UserCreate	true	"User information"
//	@Success		201			{object}	models.Response
//	@Failure		400			{object}	models.Response	"Bad Request"
//	@Failure		401			{object}	models.Response	"Auth is required"
//	@Failure		403			{object}	models.Response	"Not Authorized"
//	@Failure		500			{object}	models.Response
//	@Router			/user [post]
func CreateUser(c *fiber.Ctx) error {
	var userCreate models.UserCreate
	if err := c.BodyParser(&userCreate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Invalid request",
		})
	}
	if err := userCreate.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}
	userCreated, err := services.UserCreate(&userCreate)
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
	return c.Status(fiber.StatusCreated).JSON(models.Response{
		Status:  true,
		Body:    userCreated,
		Message: "User created",
	})
}