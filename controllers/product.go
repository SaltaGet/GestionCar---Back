package controllers

import (
	"github.com/DanielChachagua/GestionCar/models"
	"github.com/DanielChachagua/GestionCar/services"
	"github.com/gofiber/fiber/v2"
)

// ProductGetByID retrieves a product by its ID for a specific workplace.
// @Summary     Get Product By ID
// @Description Get a product or part by its ID within a specified workplace.
// @Tags        product
// @Accept      json
// @Produce     json
// @Param       id   path      string  true  "ID of the product"
// @Success     200  {object}  models.Response{body=models.ProductLaundry} "Product obtained with success"
// @Success     200  {object}  models.Response{body=models.PartWorkshop} "Part obtained with success"
// @Failure     400  {object}  models.Response "ID is required or Workplace is required"
// @Failure     500  {object}  models.Response "Internal server error"
// @Router      /product/{id} [get]
// @Security    BearerAuth

func ProductGetByID(c *fiber.Ctx) error {
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

	laundry, workshop, err := services.ProductGetByID(id, workplace.Identifier)
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
		Message: "Producto obtenido con éxito",
	})
	}

	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    workshop,
		Message: "Parte obtenida con éxito",
	})
}

// ProductGetAll godoc
// @Summary     Get All Products
// @Description Get All Products
// @Tags        product
// @Accept      json
// @Produce     json
// @Success     200 {object} models.Response{body=[]models.Laundry} "Products obtained with success"
// @Success     200 {object} models.Response{body=[]models.Workshop} "Workshop parts obtained with success"
// @Failure     400 {object} models.Response "Workplace is required"
// @Failure     500 {object} models.Response "Internal server error"
// @Router      /product [get]
// @Security    BearerAuth
func ProductGetAll(c *fiber.Ctx) error {
	workplace := c.Locals("workplace").(*models.Workplace)
	if workplace == nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Workplace is required",
		})
	}

	laundry, workshop, err := services.ProductGetAll(workplace.Identifier)
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
			Message: "Productos obtenidos con éxito",
		})
	}

	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    workshop,
		Message: "Partes obtenidas con éxito",
	})
}

// ProductGetByName retrieves products by their name from the specified workplace.
// @Summary     Get Product By Name
// @Description Fetches products from either laundry or workshop based on the provided name and workplace.
// @Tags        Product
// @Accept      json
// @Produce     json
// @Param       name  query     string  true  "Name of the Product"
// @Success     200   {object}  models.Response{body=[]models.ProductLaundry} "List of laundry products"
// @Success     200   {object}  models.Response{body=[]models.PartWorkshop} "List of workshop parts"
// @Failure     400   {object}  models.Response "Invalid name or workplace"
// @Failure     500   {object}  models.Response "Internal server error"
// @Router      /product/by-name [get]
// @Security    BearerAuth
func ProductGetByName(c *fiber.Ctx) error {
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

	laundry, workshop, err := services.ProductGetByName(name, workplace.Identifier)
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
			Message: "Productos obtenidos con éxito",
		})
	}

	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    workshop,
		Message: "Partes obtenidas con éxito",
	})
}


// ProductGetByIdentifier godoc
// @Summary     Get Products by identifier
// @Description Get Products by identifier
// @Tags        product
// @Accept      json
// @Produce     json
// @Param       identifier query string true "Identifier of product"
// @Success     200 {object} models.Response{body=[]models.Laundry} "Products obtained with success"
// @Success     200 {object} models.Response{body=[]models.Workshop} "Workshop parts obtained with success"
// @Failure     400 {object} models.Response "Workplace is required"
// @Failure     500 {object} models.Response "Internal server error"
// @Router      /product/identifier [get]
// @Security    BearerAuth
func ProductGetByIdentifier(c *fiber.Ctx) error {
	name := c.Query("identifier")
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

	laundry, workshop, err := services.ProductGetByIdentifier(name, workplace.Identifier)
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
			Message: "Productos obtenidos con éxito",
		})
	}

	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    workshop,
		Message: "Partes obtenidas con éxito",
	})
}

// ProductUpdateStock godoc
// @Summary     Update Product Stock
// @Description Updates the stock of a product based on the given method (add, subtract, update).
// @Tags        product
// @Accept      json
// @Produce     json
// @Param       id      path      string            true  "ID of the product"
// @Param       method  query     string            true  "Method of stock update (add, subtract, update)"
// @Param       stock   body      models.StockUpdate true  "Stock update details"
// @Success     200     {object}  models.Response   "Product stock updated successfully"
// @Failure     400     {object}  models.Response   "Invalid request or missing required parameters"
// @Failure     500     {object}  models.Response   "Internal server error"
// @Router      /product/{id}/stock [put]
// @Security    BearerAuth
func ProductUpdateStock(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "ID is required",
		})
	}

	method := c.Query("method")
	if method == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Method is required",
		})
	}

	var stockUpdate models.StockUpdate
	if err := c.BodyParser(&stockUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Invalid request",
		})
	}
	if err := stockUpdate.Validate(); err != nil {
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

	err := services.ProductUpdateStock(id, &stockUpdate, method, workplace.Identifier)
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
		Message: "Producto actualizado con éxito",
	})
}

// ProductUpdate godoc
// @Summary     Update Product
// @Description Updates the given product and returns the updated product.
// @Tags        product
// @Accept      json
// @Produce     json
// @Param       id      path      string            true  "ID of the product"
// @Param       product body      models.ProductUpdate true  "Product update details"
// @Success     200     {object}  models.Response   "Product updated successfully"
// @Failure     400     {object}  models.Response   "Invalid request or missing required parameters"
// @Failure     500     {object}  models.Response   "Internal server error"
// @Router      /product/{id} [put]
// @Security    BearerAuth
func ProductUpdate(c *fiber.Ctx) error {
	var productUpdate models.ProductUpdate
	if err := c.BodyParser(&productUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Invalid request",
		})
	}
	if err := productUpdate.Validate(); err != nil {
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

	err := services.ProductUpdate(&productUpdate, workplace.Identifier)
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
		Message: "Producto actualizado con éxito",
	})
}

// ProductDelete godoc
// @Summary     Delete Product
// @Description Deletes the given product with the given id.
// @Tags        product
// @Accept      json
// @Produce     json
// @Param       id   path      string     true  "ID of the product"
// @Success     200  {object}  models.Response  "Product deleted with success"
// @Failure     400  {object}  models.Response  "Invalid request or missing required parameters"
// @Failure     500  {object}  models.Response  "Internal server error"
// @Router      /product/{id} [delete]
// @Security    BearerAuth
func ProductDelete(c *fiber.Ctx) error {
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

	err := services.ProductDelete(id, workplace.Identifier)
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
		Message: "Producto eliminado con éxito",
	})
}

// ProductCreate godoc
// @Summary     Create Product
// @Description Creates a new product in the specified workplace.
// @Tags        product
// @Accept      json
// @Produce     json
// @Param       product body      models.ProductCreate true "Details of the product to create"
// @Success     200     {object}  models.Response      "Product created successfully"
// @Failure     400     {object}  models.Response      "Invalid request or validation error"
// @Failure     500     {object}  models.Response      "Internal server error"
// @Router      /product [post]
// @Security    BearerAuth
func ProductCreate(c *fiber.Ctx) error {
	var productCreate models.ProductCreate
	if err := c.BodyParser(&productCreate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Invalid request",
		})
	}
	if err := productCreate.Validate(); err != nil {
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

	productCreated, err := services.ProductCreate(&productCreate, workplace.Identifier)
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
		Body:    productCreated,
		Message: "Producto creado con éxito",
	})
}

