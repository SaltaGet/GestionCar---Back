package controllers

import (
	"github.com/DanielChachagua/GestionCar/models"
	"github.com/DanielChachagua/GestionCar/services"
	"github.com/gofiber/fiber/v2"
)

//  Login Login user
//	@Summary		Login user
//	@Description	Login user required identifier and password
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			credentials	body		models.AuthLogin	true	"Credentials"
//	@Success		200			{object}	models.Response
//	@Failure		400			{object}	models.Response
//	@Failure		401			{object}	models.Response
//	@Failure		422			{object}	models.Response
//	@Failure		404			{object}	models.Response
//	@Failure		500			{object}	models.Response
//	@Router			/auth/login [post]
func AuthLogin(c *fiber.Ctx) error {
	var loginRequest models.AuthLogin
	if err := c.BodyParser(&loginRequest); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status: false, 
			Body: nil, 
			Message: "Invalid request",
		})
	}

	if err := loginRequest.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status: false, 
			Body: nil, 
			Message: err.Error(),
		})
	}

	token, err := services.AuthLogin(loginRequest.Username, loginRequest.Password)
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

	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    token,
		Message: "Token obtenido con éxito",
	})
}

//  Login Login workplace
//	@Summary		Login Workplace
//	@Description	Login workplace required workplace_id
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			workplace_id	path		string	true	"workplace_id"
//	@Success		200				{object}	models.Response
//	@Failure		400				{object}	models.Response
//	@Failure		401				{object}	models.Response
//	@Failure		422				{object}	models.Response
//	@Failure		404				{object}	models.Response
//	@Failure		500				{object}	models.Response
//	@Router			/auth/workplace_login/{workplace_id} [get]
func AuthWorkplace(c *fiber.Ctx) error {
	id := c.Params("workplace_id")

	token, err := services.AuthWorkplace(id)
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

	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    token,
		Message: "Token obtenido con éxito",
	})
}