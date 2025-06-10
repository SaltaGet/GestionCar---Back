package controllers

import (
	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/gofiber/fiber/v2"
)

// PurchaseOrderGetByID godoc
//	@Summary		Get Purchase Order By ID
//	@Description	Retrieves a specific purchase order by its ID. Returns purchase order based on the tenant context.
//	@Tags			Purchase Order
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path		string										true	"ID of Purchase Order"
//	@Success		200	{object}	models.Response{body=models.PurchaseOrder}	"Laundry order obtained successfully"
//	@Failure		400	{object}	models.Response								"Bad Request"
//	@Failure		401	{object}	models.Response								"Auth is required"
//	@Failure		403	{object}	models.Response								"Not Authorized"
//	@Failure		404	{object}	models.Response								"Purchase Order not found"
//	@Failure		500	{object}	models.Response								"Internal server error"
//	@Router			/purchase_order/{id} [get]
func (p *PurchaseOrderController) PurchaseOrderGetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "ID is required",
		})
	}

	purchaseOrder, err := p.PurchaseOrderService.PurchaseOrderGetByID(id)
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
		Body:    purchaseOrder,
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
//	@Success		200	{object}	models.Response{body=[]models.PurchaseOrder}	"Purchase Orders obtained with success"
//	@Failure		400	{object}	models.Response									"Bad Request"
//	@Failure		401	{object}	models.Response									"Auth is required"
//	@Failure		403	{object}	models.Response									"Not Authorized"
//	@Failure		500	{object}	models.Response									"Internal server error"
//	@Router			/purchase_order/get_all [get]
//	@Security		BearerAuth
func (p *PurchaseOrderController) PurchaseOrderGetAll(c *fiber.Ctx) error {
	purchasesOrder, err := p.PurchaseOrderService.PurchaseOrderGetAll()
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
		Body:    purchasesOrder,
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
//	@Param			purchaseOrderCreate	body		models.PurchaseOrderCreate		true	"Purchase order creation data"
//	@Success		200					{object}	models.Response{body=string}	"Purchase order created successfully"
//	@Failure		400					{object}	models.Response					"Bad Request"
//	@Failure		401					{object}	models.Response					"Auth is required"
//	@Failure		403					{object}	models.Response					"Not Authorized"
//	@Failure		422					{object}	models.Response					"Model invalid"
//	@Failure		500					{object}	models.Response					"Internal server error"
//	@Router			/purchase_order/create     [post]
//	@Security		BearerAuth
func (p *PurchaseOrderController) PurchaseOrderCreate(c *fiber.Ctx) error {
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

	id, err := p.PurchaseOrderService.PurchaseOrderCreate(&purchaseOrderCreate)
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
//	@Param			purchaseOrderUpdate	body		models.PurchaseOrderUpdate	true	"Purchase order update data"
//	@Success		200					{object}	models.Response				"Purchase order updated successfully"
//	@Failure		400					{object}	models.Response				"Bad Request"
//	@Failure		401					{object}	models.Response				"Auth is required"
//	@Failure		403					{object}	models.Response				"Not Authorized"
//	@Failure		422					{object}	models.Response				"Model invalid"
//	@Failure		500					{object}	models.Response				"Internal server error"
//	@Router			/purchase_order/update [put]
func (p *PurchaseOrderController) PurchaseOrderUpdate(c *fiber.Ctx) error {
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

	err := p.PurchaseOrderService.PurchaseOrderUpdate(&purchaseOrderUpdate)
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
//	@Param			id	path		string			true	"ID of Purchase Order"
//	@Success		200	{object}	models.Response	"Purchase order deleted successfully"
//	@Failure		400	{object}	models.Response	"Bad Request"
//	@Failure		401	{object}	models.Response	"Auth is required"
//	@Failure		403	{object}	models.Response	"Not Authorized"
//	@Failure		404	{object}	models.Response	"Purchase order not found"
//	@Failure		500	{object}	models.Response	"Internal server error"
//	@Router			/purchase_order/delete/{id} [delete]
func (p *PurchaseOrderController) PurchaseOrderDelete(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "ID is required",
		})
	}

	err := p.PurchaseOrderService.PurchaseOrderDelete(id)
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

