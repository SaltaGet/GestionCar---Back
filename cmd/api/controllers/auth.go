package controllers

import (
	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/gofiber/fiber/v2"
)

//  Login godoc
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
func (a *AuthController) AuthLogin(c *fiber.Ctx) error {
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

	token, err := a.AuthService.AuthLogin(loginRequest.Username, loginRequest.Password)
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

//  LoginTenant godoc
//	@Summary		Login Tenant
//	@Description	Login tenant required tenant_id
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			tenant_id	path		string	true	"tenant_id"
//	@Success		200			{object}	models.Response
//	@Failure		400			{object}	models.Response
//	@Failure		401			{object}	models.Response
//	@Failure		403			{object}	models.Response
//	@Failure		404			{object}	models.Response
//	@Failure		422			{object}	models.Response
//	@Failure		500			{object}	models.Response
//	@Router			/auth/tenant_login/{tenant_id} [get]
func (a *AuthController) AuthTenant(c *fiber.Ctx) error {
	id := c.Params("tenant_id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status: false, 
			Body: nil, 
			Message: "ID is required",
		})
	}

	user := c.Locals("user").(*models.AuthenticatedUser)
	if user == nil {
		return c.Status(401).JSON(models.Response{
			Status: false, 
			Body: nil, 
			Message: "login is required",
		})
	}

	token, err := a.AuthService.AuthGetTenant(user, id)
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