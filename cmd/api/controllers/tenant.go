package controllers

import (
	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/gofiber/fiber/v2"
)

//  Workplace GetAll
//	@Summary		Workplace GetAll
//	@Description	Workplace GetAll required auth token
//	@Tags			Workplace
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200	{object}	models.Response{body=[]models.Workplace}	"Workplaces obtenidos con éxito"
//	@Failure		400	{object}	models.Response	"Bad Request"
//	@Failure		401	{object}	models.Response	"Auth is required"
//	@Failure		403	{object}	models.Response	"Not Authorized"
//	@Failure		500	{object}	models.Response
//	@Router			/tenant/get_all [get]
func (t *TenantController) GetWorkplaces(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	workplaces, err := t.TenantService.TenantGetByID(user.ID)
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