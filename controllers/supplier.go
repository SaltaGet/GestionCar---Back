package controllers

import (
	"github.com/DanielChachagua/GestionCar/models"
	"github.com/DanielChachagua/GestionCar/services"
	"github.com/gofiber/fiber/v2"
)

// SupplierGetByID godoc
// @Summary     Get Supplier By ID
// @Description Get a supplier by its ID within a specified workplace.
// @Tags        supplier
// @Accept      json
// @Produce     json
// @Param       id   path      string  true  "ID of the supplier"
// @Success     200  {object}  models.Response{body=models.SupplierLaundry} "Supplier obtained with success"
// @Success     200  {object}  models.Response{body=models.SupplierWorkshop} "Workshop supplier obtained with success"
// @Failure     400  {object}  models.Response "ID is required or Workplace is required"
// @Failure     500  {object}  models.Response "Internal server error"
// @Router      /supplier/{id} [get]
// @Security    BearerAuth
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
// @Summary     Get All Suppliers
// @Description Get All Suppliers
// @Tags        supplier
// @Accept      json
// @Produce     json
// @Success     200 {object} models.Response{body=[]models.SupplierLaundry} "Suppliers obtained with success"
// @Success     200 {object} models.Response{body=[]models.SupplierWorkshop} "Workshop suppliers obtained with success"
// @Failure     400 {object} models.Response "Workplace is required"
// @Failure     500 {object} models.Response "Internal server error"
// @Router      /supplier [get]
// @Security    BearerAuth
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

// SupplierGetByName retrieves suppliers by their name from the specified workplace.
// @Summary     Get Supplier By Name
// @Description Fetches suppliers from either laundry or workshop based on the provided name and workplace.
// @Tags        Supplier
// @Accept      json
// @Produce     json
// @Param       name  query     string  true  "Name of the Supplier"
// @Success     200   {object}  models.Response{body=[]models.SupplierLaundry} "List of laundry suppliers"
// @Success     200   {object}  models.Response{body=[]models.SupplierWorkshop} "List of workshop suppliers"
// @Failure     400   {object}  models.Response "Invalid name or workplace"
// @Failure     500   {object}  models.Response "Internal server error"
// @Router      /supplier/by-name [get]
// @Security    BearerAuth
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
// @Summary     Create Supplier
// @Description Creates a new supplier within the specified workplace.
// @Tags        supplier
// @Accept      json
// @Produce     json
// @Param       supplier body      models.SupplierCreate true "Details of the supplier to create"
// @Success     200     {object}  models.Response{body=string} "Supplier created successfully"
// @Failure     400     {object}  models.Response      "Invalid request or validation error"
// @Failure     500     {object}  models.Response      "Internal server error"
// @Router      /supplier [post]
// @Security    BearerAuth
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

// SupplierUpdate godoc
// @Summary     Update Supplier
// @Description Update a supplier's information from the specified workplace.
// @Tags        supplier
// @Accept      json
// @Produce     json
// @Param       id   path      string  true  "ID of the Supplier"
// @Param       body body      models.SupplierUpdate  true  "Supplier information"
// @Success     200  {object}  models.Response "Supplier updated with success"
// @Failure     400  {object}  models.Response "Invalid request or Validation error"
// @Failure     500  {object}  models.Response "Internal server error"
// @Router      /supplier/{id} [put]
// @Security    BearerAuth
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

// SupplierDeleteByID deletes a supplier by its ID for a specific workplace.
// @Summary     Delete Supplier
// @Description Deletes a supplier based on the provided ID and workplace context.
// @Tags        supplier
// @Accept      json
// @Produce     json
// @Param       id   path      string  true  "ID of the supplier"
// @Success     200  {object}  models.Response  "Supplier deleted with success"
// @Failure     400  {object}  models.Response  "ID is required or Workplace is required"
// @Failure     500  {object}  models.Response  "Internal server error"
// @Router      /supplier/{id} [delete]
// @Security    BearerAuth
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

