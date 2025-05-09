package controllers

import (
	"github.com/DanielChachagua/GestionCar/models"
	"github.com/DanielChachagua/GestionCar/services"
	"github.com/gofiber/fiber/v2"
)

//  Workplace GetAll
//  @Summary		Workplace GetAll
//	@Description	Workplace GetAll required auth token
//	@Tags			Workplace
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200			{object}	models.Response
//	@Failure		400			{object}	models.Response
//	@Failure		401			{object}	models.Response
//	@Failure		422			{object}	models.Response
//	@Failure		404			{object}	models.Response
//	@Failure		500			{object}	models.Response
//	@Router			/workplace/get_all [get]
func GetWorkplaces(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	workplaces, err := services.GetWorkplaceAll(user.Role)
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
		Body:    workplaces,
		Message: "Workplaces obtenidos con éxito",
	})
}