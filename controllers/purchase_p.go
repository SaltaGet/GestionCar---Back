package controllers

import (
	"github.com/DanielChachagua/GestionCar/models"
	"github.com/DanielChachagua/GestionCar/services"
	"github.com/gofiber/fiber/v2"
)
// GetEmployeeByID godoc
// @Summary     Get Employee By ID
// @Description Get Employee By ID
// @Tags        employee
// @Accept      json
// @Produce     json
// @Param       id   path      string  true  "ID of Employee"
// @Success     200  {object}  models.Response
// @Failure     400  {object}  models.Response
// @Failure     404  {object}  models.Response
// @Failure     500  {object}  models.Response
// @Router      /employee/{id} [get]
// @Security    BearerAuth
func PurchaseProductGetByID(c *fiber.Ctx) error {
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

	laundry, workshop, err := services.PurchaseProductGetByID(id, workplace.Identifier)
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
			Message: "Producto de compra obtenida con éxito",
		})
	}

	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    workshop,
		Message: "Producto de compra obtenida con éxito",
	})
}

// PurchaseProductGetAllByPurhcaseID godoc
// @Summary     Get All Products From Purchase Order
// @Description Get All Products From Purchase Order
// @Tags        purchase_product
// @Accept      json
// @Produce     json
// @Param       purchase_id   path      string  true  "ID of Purchase Order"
// @Success     200           {object}  models.Response{body=[]models.PurchaseProductLaundry} "Products obtained with success"
// @Success     200           {object}  models.Response{body=[]models.PurchasePartWorkshop} "Workshop parts obtained with success"
// @Failure     400           {object}  models.Response "ID is required"
// @Failure     400           {object}  models.Response "Workplace is required"
// @Failure     404           {object}  models.Response "Purchase order not found"
// @Failure     500           {object}  models.Response "Internal server error"
// @Router      /purchase_product/{purchase_id} [get]
// @Security    BearerAuth
func PurchaseProductGetAllByPurhcaseID(c *fiber.Ctx) error {
	purchaseId := c.Params("purchase_id")
	if purchaseId == "" {
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

	laundry, workshop, err := services.PurchaseProductGetAllByPurhcaseID(purchaseId, workplace.Identifier)
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
			Message: "Productos de orden de compra obtenida con éxito",
		})
	}

	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    workshop,
		Message: "Productos de orden de compra obtenida con éxito",
	})
}

// PurchaseProductCreate godoc
// @Summary     Create Purchase Product
// @Description Creates a purchase product, either for laundry or workshop.
//              Returns the ID of the created purchase product.
// @Tags        purchase_product
// @Accept      json
// @Produce     json
// @Param       purchaseProductCreate body     models.PurchaseProductCreate true  "Purchase product creation data"
// @Success     200                 {object} models.Response{body=string} "Purchase product created successfully"
// @Failure     400                 {object} models.Response            "Invalid request"
// @Failure     400                 {object} models.Response            "Workplace is required"
// @Failure     500                 {object} models.Response            "Internal server error"
// @Router      /purchase_product   [post]
// @Security    BearerAuth
func PurchaseProductCreate(c *fiber.Ctx) error {
	var purchaseProductCreate models.PurchaseProductCreate
	if err := c.BodyParser(&purchaseProductCreate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Invalid request",
		})
	}
	if err := purchaseProductCreate.Validate(); err != nil {
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

	id, err := services.PurchaseProductCreate(&purchaseProductCreate, workplace.Identifier)
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
		Message: "Producto deOrden de compra creado con éxito",
	})
}

// PurchaseProductUpdate godoc
// @Summary     Update Purchase Product
// @Description Updates the given purchase product and returns a success message.
// @Tags        purchase_product
// @Accept      json
// @Produce     json
// @Param       id      path      string            true  "ID of the purchase product"
// @Param       product body      models.PurchaseProductUpdate true  "Purchase product update details"
// @Success     200     {object}  models.Response   "Purchase product updated successfully"
// @Failure     400     {object}  models.Response   "Invalid request or missing required parameters"
// @Failure     500     {object}  models.Response   "Internal server error"
// @Router      /purchase_product/{id} [put]
// @Security    BearerAuth
func PurchaseProductUpdate(c *fiber.Ctx) error {
	var purchaseProductUpdate models.PurchaseProductUpdate
	if err := c.BodyParser(&purchaseProductUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Invalid request",
		})
	}
	if err := purchaseProductUpdate.Validate(); err != nil {
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

	err := services.PurchaseProductUpdate(&purchaseProductUpdate, workplace.Identifier)
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
		Message: "Producto de Orden de compra editado con éxito",
	})
}

// PurchaseProductDelete godoc
// @Summary     Delete Purchase Product
// @Description Deletes a specific purchase product.
//              Returns a success message if the deletion is successful.
// @Tags        purchase_product
// @Accept      json
// @Produce     json
// @Param       id   path      string  true  "ID of Purchase Product"
// @Success     200  {object}  models.Response "Purchase product deleted successfully"
// @Failure     400  {object}  models.Response "ID is required"
// @Failure     400  {object}  models.Response "Workplace is required"
// @Failure     500  {object}  models.Response "Internal server error"
// @Router      /purchase_product/{id} [delete]
// @Security    BearerAuth
func PurchaseProductDelete(c *fiber.Ctx) error {
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

	err := services.PurchaseProductDelete(id, workplace.Identifier)
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
		Message: "Producto deOrden de compra eliminado con éxito",
	})
}
