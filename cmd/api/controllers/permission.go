package controllers

import (
	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/gofiber/fiber/v2"
)

//	Permissions GetAll godoc
//
// @Summary		Permissions GetAll
// @Description	Permissions GetAll required auth token
// @Tags			Member
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Success		200				{object}	models.Response		"Members obtenidos con éxito"
// @Failure		400				{object}	models.Response		"Bad Request"
// @Failure		401				{object}	models.Response		"Auth is required"
// @Failure		403				{object}	models.Response		"Not Authorized"
// @Failure		500				{object}	models.Response
// @Router			/permission/get_all [post]
func (p *PermissionController) PermissionGetAll(c *fiber.Ctx) error {
	permissions, err := p.PermissionService.PermissionGetAll()
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

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Status:  true,
		Body:    permissions,
		Message: "Permisos obtenidos con éxito",
	})
}

//	Permissions GetAllToMe godoc
// @Summary		Permissions GetAlltoME
// @Description	Permissions GetAllToMe required auth token
// @Tags			Member
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Success		200				{object}	models.Response		"Members obtenidos con éxito"
// @Failure		400				{object}	models.Response		"Bad Request"
// @Failure		401				{object}	models.Response		"Auth is required"
// @Failure		403				{object}	models.Response		"Not Authorized"
// @Failure		500				{object}	models.Response
// @Router			/permission/get_to_me [post]
func (p *PermissionController) PermissionGetToMe(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.AuthenticatedUser)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Usuario requerido",
		})
	}

	permissions, err := p.PermissionService.PermissionGetToMe(*user.RoleID)
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

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Status:  true,
		Body:    permissions,
		Message: "Permisos obtenidos con éxito",
	})
}