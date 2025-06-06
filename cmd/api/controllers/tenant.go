package controllers

import (
	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/gofiber/fiber/v2"
)

//	 Tenant godoc
//		@Summary		Tenant GetAll
//		@Description	Tenant GetAll required auth token
//		@Tags			Tenant
//		@Accept			json
//		@Produce		json
//		@Security		BearerAuth
//		@Success		200	{object}	models.Response{body=[]models.TenantResponse}	"Tenants obtenidos con éxito"
//		@Failure		400	{object}	models.Response									"Bad Request"
//		@Failure		401	{object}	models.Response									"Auth is required"
//		@Failure		403	{object}	models.Response									"Not Authorized"
//		@Failure		500	{object}	models.Response
//		@Router			/tenant/get_all [get]
func (t *TenantController) GetTenants(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	tenants, err := t.TenantService.TenantGetAll(user.ID)
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

	if tenants == nil || len(*tenants) == 0 {
		empty := []models.TenantResponse{}
		tenants = &empty
	}

	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    *tenants,
		Message: "Tenants obtenidos con éxito",
	})
}

//	 Tenant godoc
//		@Summary		Tenant Create
//		@Description	Tenant Create required auth token
//		@Tags			Tenant
//		@Accept			json
//		@Produce		json
//		@Security		BearerAuth
//		@Param			user_id	query		string	true	"UserID"
//		@Param			TenantCreate	body		models.TenantCreate	true	"TenantCreate"
//		@Success		200				{object}	models.Response		"Tenant creado con éxito"
//		@Failure		400				{object}	models.Response		"Bad Request"
//		@Failure		401				{object}	models.Response		"Auth is required"
//		@Failure		403				{object}	models.Response		"Not Authorized"
//		@Failure		500				{object}	models.Response
//		@Router			/tenant/create [post]
func (t *TenantController) TenantCreateByUserID(c *fiber.Ctx) error {
	userID := c.Query("user_id")
	if userID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "El user_id no debe de ser vacio",
		})
	}

	var tenantCreate models.TenantCreate
	if err := c.BodyParser(&tenantCreate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Invalid request",
		})
	}
	if err := tenantCreate.Validate(); err != nil {
		return c.Status(422).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	id, err := t.TenantService.TenantCreateByUserID(&tenantCreate, userID)
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
		Body:    id,
		Message: "Tenant creado con éxito",
	})
}

//	 Tenant godoc
//		@Summary		Tenant Create
//		@Description	Tenant Create required auth token
//		@Tags			Tenant
//		@Accept			json
//		@Produce		json
//		@Security		BearerAuth
//		@Param			TenantUserCreate	body		models.TenantUserCreate	true	"TenantUserCreate"
//		@Success		200				{object}	models.Response		"Tenant y Usuario creados con éxito"
//		@Failure		400				{object}	models.Response		"Bad Request"
//		@Failure		401				{object}	models.Response		"Auth is required"
//		@Failure		403				{object}	models.Response		"Not Authorized"
//		@Failure		500				{object}	models.Response
//		@Router			/tenant/create_tenant_user [post]
func (t *TenantController) TenantUserCreate(c *fiber.Ctx) error {
	var tenantUserCrate models.TenantUserCreate
	if err := c.BodyParser(&tenantUserCrate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Invalid request",
		})
	}
	if err := tenantUserCrate.TenantCreate.Validate(); err != nil {
		return c.Status(422).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}
	if err := tenantUserCrate.UserCreate.Validate(); err != nil {
		return c.Status(422).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	id, err := t.TenantService.TenantUserCreate(&tenantUserCrate)
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
		Body:    id,
		Message: "Tenant creado con éxito",
	})
}
