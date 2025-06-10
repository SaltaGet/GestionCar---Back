package controllers

import (
	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/gofiber/fiber/v2"
)

//	Member godoc
//
// @Summary		Memeber GetAll
// @Description	Memeber GetAll required auth token
// @Tags			Member
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Success		200	{object}	models.Response{body=[]models.Member}	"Members obtenidos con éxito"
// @Failure		400	{object}	models.Response							"Bad Request"
// @Failure		401	{object}	models.Response							"Auth is required"
// @Failure		403	{object}	models.Response							"Not Authorized"
// @Failure		500	{object}	models.Response
// @Router			/member/get_all [get]
func (m *MemberController) MemberGetAll(c *fiber.Ctx) error {
	memebers, err := m.MemberService.MemberGetAll()
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
		Body:    memebers,
		Message: "Memebers received successfully",
	})
}

//	Member godoc
//
// @Summary		Memeber GetAll
// @Description	Memeber GetAll required auth token
// @Tags			Member
// @Accept			json
// @Produce		json
// @Security		BearerAuth
// @Param			member_create	body		models.MemberCreate	true	"MemberCreate"
// @Success		200				{object}	models.Response		"Members obtenidos con éxito"
// @Failure		400				{object}	models.Response		"Bad Request"
// @Failure		401				{object}	models.Response		"Auth is required"
// @Failure		403				{object}	models.Response		"Not Authorized"
// @Failure		500				{object}	models.Response
// @Router			/member/create [post]
func (m *MemberController) MemberCreate(c *fiber.Ctx) error {
	user, ok := c.Locals("user").(*models.AuthenticatedUser)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Teanant No autenticado",
		})
	}

	var memberCreate models.MemberCreate
	if err := c.BodyParser(&memberCreate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Invalid request",
		})
	}
	if err := memberCreate.Validate(); err != nil {
		return c.Status(422).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	id, err := m.MemberService.MemberCreate(&memberCreate, user)
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
		Body:    id,
		Message: "Memebers received successfully",
	})
}
