package controllers

import (
	"github.com/DanielChachagua/GestionCar/models"
	"github.com/DanielChachagua/GestionCar/services"
	"github.com/gofiber/fiber/v2"
)

// PurchaseOrderGetByID godoc
//	@Summary		Get Purchase Order By ID
//	@Description	Retrieves a specific purchase order by its ID.
//              Returns either a laundry or workshop purchase order based on the workplace context.
//	@Tags			Purchase Order
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			X-Workplace-Token	header		string												true	"Workplace Token"
//	@Param			id					path		string												true	"ID of Purchase Order"
//	@Success		200					{object}	models.Response{body=models.PurchaseOrderLaundry}	"Laundry order obtained successfully"
//	@Failure		400					{object}	models.Response										"Bad Request"
//	@Failure		401					{object}	models.Response										"Auth is required"
//	@Failure		403					{object}	models.Response										"Not Authorized"
//	@Failure		404					{object}	models.Response										"Purchase Order not found"
//	@Failure		500					{object}	models.Response										"Internal server error"
//	@Router			/purchase_order/{id} [get]
func PurchaseOrderGetByID(c *fiber.Ctx) error {
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

	laundry, workshop, err := services.PurchaseOrderGetByID(id, workplace.Identifier)
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
			Message: "Orden de compra obtenida con éxito",
		})
	}

	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    workshop,
		Message: "Orden de compra obtenida con éxito",
	})
}

// PurchaseOrderGetAll godoc
//	@Summary		Get All Purchase Orders
//	@Description	Get All Purchase Orders
//	@Tags			Purchase Order
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			X-Workplace-Token	header		string												true	"Workplace Token"
//	@Success		200					{object}	models.Response{body=[]models.PurchaseOrderLaundry}	"Purchase Orders obtained with success"
//	@Failure		400					{object}	models.Response										"Bad Request"
//	@Failure		401					{object}	models.Response										"Auth is required"
//	@Failure		403					{object}	models.Response										"Not Authorized"
//	@Failure		500					{object}	models.Response										"Internal server error"
//	@Router			/purchase_order/get_all [get]
//	@Security		BearerAuth
func PurchaseOrderGetAll(c *fiber.Ctx) error {
	workplace := c.Locals("workplace").(*models.Workplace)
	if workplace == nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Workplace is required",
		})
	}

	laundry, workshop, err := services.PurchaseOrderGetAll(workplace.Identifier)
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
			Message: "Orden de compra obtenida con éxito",
		})
	}

	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    workshop,
		Message: "Orden de compra obtenida con éxito",
	})
}

// PurchaseOrderCreate godoc
//	@Summary		Create Purchase Order
//	@Description	Creates a purchase order, either for laundry or workshop.
//	@Tags			Purchase Order
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			X-Workplace-Token	header		string							true	"Workplace Token"
//	@Param			purchaseOrderCreate	body		models.PurchaseOrderCreate		true	"Purchase order creation data"
//	@Success		200					{object}	models.Response{body=string}	"Purchase order created successfully"
//	@Failure		400					{object}	models.Response					"Bad Request"
//	@Failure		401					{object}	models.Response					"Auth is required"
//	@Failure		403					{object}	models.Response					"Not Authorized"
//	@Failure		422					{object}	models.Response					"Model invalid"
//	@Failure		500					{object}	models.Response					"Internal server error"
//	@Router			/purchase_order/create     [post]
//	@Security		BearerAuth
func PurchaseOrderCreate(c *fiber.Ctx) error {
	var purchaseOrderCreate models.PurchaseOrderCreate
	if err := c.BodyParser(&purchaseOrderCreate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Invalid request",
		})
	}
	if err := purchaseOrderCreate.Validate(); err != nil {
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

	id, err := services.PurchaseOrderCreate(&purchaseOrderCreate, workplace.Identifier)
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
		Message: "Orden de compra creada con éxito",
	})
}

// PurchaseOrderUpdate godoc
//	@Summary		Update Purchase Order
//	@Description	Updates an existing purchase order with new details.
//              Validates the request body and workplace context.
//              Returns a success message if the update is successful.
//	@Tags			Purchase Order
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			X-Workplace-Token	header		string						true	"Workplace Token"
//	@Param			purchaseOrderUpdate	body		models.PurchaseOrderUpdate	true	"Purchase order update data"
//	@Success		200					{object}	models.Response				"Purchase order updated successfully"
//	@Failure		400					{object}	models.Response				"Bad Request"
//	@Failure		401					{object}	models.Response				"Auth is required"
//	@Failure		403					{object}	models.Response				"Not Authorized"
//	@Failure		422					{object}	models.Response				"Model invalid"
//	@Failure		500					{object}	models.Response				"Internal server error"
//	@Router			/purchase_order/update [put]
func PurchaseOrderUpdate(c *fiber.Ctx) error {
	var purchaseOrderUpdate models.PurchaseOrderUpdate
	if err := c.BodyParser(&purchaseOrderUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Invalid request",
		})
	}
	if err := purchaseOrderUpdate.Validate(); err != nil {
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

	err := services.PurchaseOrderUpdate(&purchaseOrderUpdate, workplace.Identifier)
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
		Message: "Orden de compra editada con éxito",
	})
}

// PurchaseOrderDelete godoc
//	@Summary		Delete Purchase Order
//	@Description	Deletes a specific purchase order by its ID.
//	@Tags			Purchase Order
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			X-Workplace-Token	header		string			true	"Workplace Token"
//	@Param			id					path		string			true	"ID of Purchase Order"
//	@Success		200					{object}	models.Response	"Purchase order deleted successfully"
//	@Failure		400					{object}	models.Response	"Bad Request"
//	@Failure		401					{object}	models.Response	"Auth is required"
//	@Failure		403					{object}	models.Response	"Not Authorized"
//	@Failure		404					{object}	models.Response	"Purchase order not found"
//	@Failure		500					{object}	models.Response	"Internal server error"
//	@Router			/purchase_order/delete/{id} [delete]
func PurchaseOrderDelete(c *fiber.Ctx) error {
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

	err := services.PurchaseOrderDelete(id, workplace.Identifier)
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
		Message: "Orden de compra eliminada con éxito",
	})
}

