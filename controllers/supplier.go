package controllers

import (
	"github.com/DanielChachagua/GestionCar/models"
	"github.com/DanielChachagua/GestionCar/services"
	"github.com/gofiber/fiber/v2"
)

// SupplierGetByID godoc
//	@Summary		Get Supplier By ID
//	@Description	Get a supplier by its ID within a specified workplace.
//	@Tags			Supplier
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			X-Workplace-Token	header		string											true	"Workplace Token"
//	@Param			id					path		string											true	"ID of the supplier"
//	@Success		200					{object}	models.Response{body=models.SupplierLaundry}	"Supplier obtained with success"
//	@Failure		400					{object}	models.Response									"Bad Request"
//	@Failure		401					{object}	models.Response									"Auth is required"
//	@Failure		403					{object}	models.Response									"Not Authorized"
//	@Failure		404					{object}	models.Response									"Supplier not found"
//	@Failure		500					{object}	models.Response									"Internal server error"
//	@Router			/supplier/{id} [get]
func SupplierGetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "ID is required",
		})
	}

	workplace := c.Locals("workplace").(*models.Workplace)
	if workplace == nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Workplace is required",
		})
	}

	laundry, workshop, err := services.SupplierGetByID(id, workplace.Identifier)
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

	if laundry != nil {
		return c.Status(200).JSON(models.Response{
			Status:  true,
			Body:    laundry,
			Message: "Proveedor obtenido con éxito",
		})
	}

	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    workshop,
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
//	@Param			X-Workplace-Token	header		string											true	"Workplace Token"
//	@Success		200					{object}	models.Response{body=[]models.SupplierLaundry}	"Suppliers obtained with success"
//	@Failure		400					{object}	models.Response									"Bad Request"
//	@Failure		401					{object}	models.Response									"Auth is required"
//	@Failure		403					{object}	models.Response									"Not Authorized"
//	@Failure		500					{object}	models.Response									"Internal server error"
//	@Router			/supplier/get_all [get]
func SupplierGetAll(c *fiber.Ctx) error {
	workplace := c.Locals("workplace").(*models.Workplace)
	if workplace == nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Workplace is required",
		})
	}

	laundry, workshop, err := services.SupplierGetAll(workplace.Identifier)
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

	if laundry != nil {
		return c.Status(200).JSON(models.Response{
			Status:  true,
			Body:    laundry,
			Message: "Proveedores obtenidos con éxito",
		})
	}

	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    workshop,
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
//	@Param			X-Workplace-Token	header		string											true	"Workplace Token"
//	@Param			name				query		string											true	"Name of the Supplier"
//	@Success		200					{object}	models.Response{body=[]models.SupplierLaundry}	"List of suppliers"
//	@Failure		400					{object}	models.Response									"Bad Request"
//	@Failure		401					{object}	models.Response									"Auth is required"
//	@Failure		403					{object}	models.Response									"Not Authorized"
//	@Failure		500					{object}	models.Response									"Internal server error"
//	@Router			/supplier/get_by_name [get]
//	@Security		BearerAuth
func SupplierGetByName(c *fiber.Ctx) error {
	name := c.Query("name")
	if name == "" || len(name) < 3 {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "El valor no debe de ser vacio o menor a 3 caracteres",
		})
	}

	workplace := c.Locals("workplace").(*models.Workplace)
	if workplace == nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Workplace is required",
		})
	}

	laundry, workshop, err := services.SupplierGetByName(name, workplace.Identifier)
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

	if laundry != nil {
		return c.Status(200).JSON(models.Response{
			Status:  true,
			Body:    laundry,
			Message: "Proveedores obtenidos con éxito",
		})
	}

	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    workshop,
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
//	@Param			X-Workplace-Token	header		string							true	"Workplace Token"
//	@Param			supplier			body		models.SupplierCreate			true	"Details of the supplier to create"
//	@Success		200					{object}	models.Response{body=string}	"Supplier created successfully"
//	@Failure		400					{object}	models.Response					"Bad Request"
//	@Failure		401					{object}	models.Response					"Auth is required"
//	@Failure		403					{object}	models.Response					"Not Authorized"
//	@Failure		422					{object}	models.Response					"Model is invalid"
//	@Failure		500					{object}	models.Response					"Internal server error"
//	@Router			/supplier/create [post]
func SupplierCreate(c *fiber.Ctx) error {
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

	workplace := c.Locals("workplace").(*models.Workplace)
	if workplace == nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Workplace is required",
		})
	}

	id, err := services.SupplierCreate(&supplierCreate, workplace.Identifier)
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
// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6IjBlYzIzOGU5LThkZTEtNDg2MS05OGY0LTc5NjY4ZmUzZjZhNCJ9.iT0dBWAfpeFtHSYzuh7MxZrQzZ6XVzypKsreFVXksnw
// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6Ijg3OGZkNWRhLTdiNDctNDI2ZS1iZmRhLTFiMjMwMTBlMWJhNCJ9.yGhFvYKi6Zm2HpxNravhZUzE2stcr6AwrWWsNI5W-o8
// SupplierUpdate godoc
//	@Summary		Update Supplier
//	@Description	Update a supplier's information from the specified workplace.
//	@Tags			Supplier
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			X-Workplace-Token	header		string					true	"Workplace Token"
//	@Param			body				body		models.SupplierUpdate	true	"Supplier information"
//	@Success		200					{object}	models.Response			"Supplier updated with success"
//	@Failure		400					{object}	models.Response			"Bad Request"
//	@Failure		401					{object}	models.Response			"Auth is required"
//	@Failure		403					{object}	models.Response			"Not Authorized"
//	@Failure		404					{object}	models.Response			"Supplier not found"
//	@Failure		422					{object}	models.Response			"Model is invalid"
//	@Failure		500					{object}	models.Response			"Internal server error"
//	@Router			/supplier/update [put]
func SupplierUpdate(c *fiber.Ctx) error {
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

	workplace := c.Locals("workplace").(*models.Workplace)
	if workplace == nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Workplace is required",
		})
	}

	err := services.SupplierUpdate(&supplierUpdate, workplace.Identifier)
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
//	@Param			X-Workplace-Token	header		string			true	"Workplace Token"
//	@Param			id					path		string			true	"ID of the supplier"
//	@Success		200					{object}	models.Response	"Supplier deleted with success"
//	@Failure		400					{object}	models.Response	"Bad Request"
//	@Failure		401					{object}	models.Response	"Auth is required"
//	@Failure		403					{object}	models.Response	"Not Authorized"
//	@Failure		404					{object}	models.Response	"Supplier not found"
//	@Failure		500					{object}	models.Response	"Internal server error"
//	@Router			/supplier/delete/{id} [delete]
func SupplierDeleteByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "ID is required",
		})
	}

	workplace := c.Locals("workplace").(*models.Workplace)
	if workplace == nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Workplace is required",
		})
	}

	err := services.SupplierDeleteByID(id, workplace.Identifier)
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

