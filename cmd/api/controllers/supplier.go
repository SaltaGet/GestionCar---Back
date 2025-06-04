package controllers

import (
	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/gofiber/fiber/v2"
)

// SupplierGetByID godoc
//	@Summary		Get Supplier By ID
//	@Description	Get a supplier by its ID within a specified workplace.
//	@Tags			Supplier
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id					path		string											true	"ID of the supplier"
//	@Success		200					{object}	models.Response{body=models.SupplierLaundry}	"Supplier obtained with success"
//	@Failure		400					{object}	models.Response									"Bad Request"
//	@Failure		401					{object}	models.Response									"Auth is required"
//	@Failure		403					{object}	models.Response									"Not Authorized"
//	@Failure		404					{object}	models.Response									"Supplier not found"
//	@Failure		500					{object}	models.Response									"Internal server error"
//	@Router			/supplier/{id} [get]
func (s *SupplierController) SupplierGetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "ID is required",
		})
	}

	supplier, err := s.SupplierService.SupplierGetByID(id)
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
		Body:    supplier,
		Message: "Proveedor obtenido con éxito",
	})
}

// SupplierGetAll godoc
//	@Summary		Get All Suppliers
//	@Description	Get All Suppliers
//	@Tags			Supplier
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200					{object}	models.Response{body=[]models.SupplierLaundry}	"Suppliers obtained with success"
//	@Failure		400					{object}	models.Response									"Bad Request"
//	@Failure		401					{object}	models.Response									"Auth is required"
//	@Failure		403					{object}	models.Response									"Not Authorized"
//	@Failure		500					{object}	models.Response									"Internal server error"
//	@Router			/supplier/get_all [get]
func (s *SupplierController) SupplierGetAll(c *fiber.Ctx) error {
	suppliers, err := s.SupplierService.SupplierGetAll()
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
		Body:    suppliers,
		Message: "Proveedores obtenidos con éxito",
	})
}

// SupplierGetByName godoc
//	@Summary		Get Supplier By Name
//	@Description	Fetches suppliers from either laundry or workshop based on the provided name and workplace.
//	@Tags			Supplier
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			name				query		string											true	"Name of the Supplier"
//	@Success		200					{object}	models.Response{body=[]models.SupplierLaundry}	"List of suppliers"
//	@Failure		400					{object}	models.Response									"Bad Request"
//	@Failure		401					{object}	models.Response									"Auth is required"
//	@Failure		403					{object}	models.Response									"Not Authorized"
//	@Failure		500					{object}	models.Response									"Internal server error"
//	@Router			/supplier/get_by_name [get]
//	@Security		BearerAuth
func (s *SupplierController) SupplierGetByName(c *fiber.Ctx) error {
	name := c.Query("name")
	if name == "" || len(name) < 3 {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "El valor no debe de ser vacio o menor a 3 caracteres",
		})
	}

	supplies, err := s.SupplierService.SupplierGetByName(name)
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
		Body:    supplies,
		Message: "Proveedores obtenidos con éxito",
	})
}

// SupplierCreate godoc
//	@Summary		Create Supplier
//	@Description	Creates a new supplier within the specified workplace.
//	@Tags			Supplier
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			supplier			body		models.SupplierCreate			true	"Details of the supplier to create"
//	@Success		200					{object}	models.Response{body=string}	"Supplier created successfully"
//	@Failure		400					{object}	models.Response					"Bad Request"
//	@Failure		401					{object}	models.Response					"Auth is required"
//	@Failure		403					{object}	models.Response					"Not Authorized"
//	@Failure		422					{object}	models.Response					"Model is invalid"
//	@Failure		500					{object}	models.Response					"Internal server error"
//	@Router			/supplier/create [post]
func (s *SupplierController) SupplierCreate(c *fiber.Ctx) error {
	var supplierCreate models.SupplierCreate
	if err := c.BodyParser(&supplierCreate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Invalid request",
		})
	}
	if err := supplierCreate.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	id, err := s.SupplierService.SupplierCreate(&supplierCreate)
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
		Message: "Proveedor creado con éxito",
	})
}

// SupplierUpdate godoc
//	@Summary		Update Supplier
//	@Description	Update a supplier's information from the specified workplace.
//	@Tags			Supplier
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			body				body		models.SupplierUpdate	true	"Supplier information"
//	@Success		200					{object}	models.Response			"Supplier updated with success"
//	@Failure		400					{object}	models.Response			"Bad Request"
//	@Failure		401					{object}	models.Response			"Auth is required"
//	@Failure		403					{object}	models.Response			"Not Authorized"
//	@Failure		404					{object}	models.Response			"Supplier not found"
//	@Failure		422					{object}	models.Response			"Model is invalid"
//	@Failure		500					{object}	models.Response			"Internal server error"
//	@Router			/supplier/update [put]
func (s *SupplierController) SupplierUpdate(c *fiber.Ctx) error {
	var supplierUpdate models.SupplierUpdate
	if err := c.BodyParser(&supplierUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Invalid request",
		})
	}
	if err := supplierUpdate.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	err := s.SupplierService.SupplierUpdate(&supplierUpdate)
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
		Body:    nil,
		Message: "Proveedor editado con éxito",
	})
}

// SupplierDeleteByID godoc
//	@Summary		Delete Supplier
//	@Description	Deletes a supplier based on the provided ID and workplace context.
//	@Tags			Supplier
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id					path		string			true	"ID of the supplier"
//	@Success		200					{object}	models.Response	"Supplier deleted with success"
//	@Failure		400					{object}	models.Response	"Bad Request"
//	@Failure		401					{object}	models.Response	"Auth is required"
//	@Failure		403					{object}	models.Response	"Not Authorized"
//	@Failure		404					{object}	models.Response	"Supplier not found"
//	@Failure		500					{object}	models.Response	"Internal server error"
//	@Router			/supplier/delete/{id} [delete]
func (s *SupplierController) SupplierDeleteByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "ID is required",
		})
	}

	err := s.SupplierService.SupplierDeleteByID(id)
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
		Body:    nil,
		Message: "Proveedor eliminado con éxito",
	})
}

