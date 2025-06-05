package controllers

import (
	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/gofiber/fiber/v2"
)

// ProductGetByID godoc
//	@Summary		Get Product By ID
//	@Description	Get a product or part by its ID within a specified workplace.
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id					path		string										true	"ID of the product"
//	@Success		200					{object}	models.Response{body=models.Product}	"Product obtained with success"
//	@Failure		400					{object}	models.Response								"Bad Request"
//	@Failure		401					{object}	models.Response								"Auth is required"
//	@Failure		403					{object}	models.Response								"Not Authorized"
//	@Failure		404					{object}	models.Response								"Expense not found"
//	@Failure		500					{object}	models.Response								"Internal server error"
//	@Router			/product/{id} [get]
func (p *ProductController) ProductGetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "ID is required",
		})
	}

	product, err := p.ProductService.ProductGetByID(id)
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
		Body:    product,
		Message: "Parte obtenida con éxito",
	})
}

// ProductGetAll godoc
//	@Summary		Get All Products
//	@Description	Get All Products
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200					{object}	models.Response{body=[]models.Product}	"Products obtained with success"
//	@Failure		400					{object}	models.Response									"Bad Request"
//	@Failure		401					{object}	models.Response									"Auth is required"
//	@Failure		403					{object}	models.Response									"Not Authorized"
//	@Failure		500					{object}	models.Response									"Internal server error"
//	@Router			/product/get_all [get]
func (p *ProductController) ProductGetAll(c *fiber.Ctx) error {
	products, err := p.ProductService.ProductGetAll()
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
		Body:    products,
		Message: "Partes obtenidas con éxito",
	})
}

// ProductGetByName godoc
//	@Summary		Get Product By Name
//	@Description	Fetches products from either laundry or workshop based on the provided name and workplace.
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			name				query		string											true	"Name of the Product"
//	@Success		200					{object}	models.Response{body=[]models.Product}	"List of products"
//	@Failure		400					{object}	models.Response									"Bad Request"
//	@Failure		401					{object}	models.Response									"Auth is required"
//	@Failure		403					{object}	models.Response									"Not Authorized"
//	@Failure		500					{object}	models.Response									"Internal server error"
//	@Router			/product/get_by_name [get]
func (p *ProductController) ProductGetByName(c *fiber.Ctx) error {
	name := c.Query("name")
	if name == "" || len(name) < 3 {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "El valor no debe de ser vacio o menor a 3 caracteres",
		})
	}

	products, err := p.ProductService.ProductGetByName(name)
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
		Body:    products,
		Message: "Partes obtenidas con éxito",
	})
}

// ProductGetByIdentifier godoc
//	@Summary		Get Products by identifier
//	@Description	Get Products by identifier
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			identifier			query		string											true	"Identifier of product"
//	@Success		200					{object}	models.Response{body=[]models.Product}	"Products obtained with success"
//	@Failure		400					{object}	models.Response									"Bad Request"
//	@Failure		401					{object}	models.Response									"Auth is required"
//	@Failure		403					{object}	models.Response									"Not Authorized"
//	@Failure		500					{object}	models.Response									"Internal server error"
//	@Router			/product/get_by_identifier [get]
func (p *ProductController) ProductGetByIdentifier(c *fiber.Ctx) error {
	identifire := c.Query("identifier")
	if identifire == "" || len(identifire) < 3 {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "El valor no debe de ser vacio o menor a 3 caracteres",
		})
	}

	products, err := p.ProductService.ProductGetByIdentifier(identifire)
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
		Body:    products,
		Message: "Partes obtenidas con éxito",
	})
}

// ProductUpdateStock godoc
//	@Summary		Update Product Stock
//	@Description	Updates the stock of a product based on the given method (add, subtract, update).
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id					path		string				true	"ID of the product"
//	@Param			method				query		string				true	"Method of stock update (add, subtract, update)"
//	@Param			stock				body		models.StockUpdate	true	"Stock update details"
//	@Success		200					{object}	models.Response		"Product stock updated successfully"
//	@Failure		400					{object}	models.Response		"Bad Request"
//	@Failure		401					{object}	models.Response		"Auth is required"
//	@Failure		403					{object}	models.Response		"Not Authorized"
//	@Failure		404					{object}	models.Response		"Product not found"
//	@Failure		422					{object}	models.Response		"Model invalid"
//	@Failure		500					{object}	models.Response		"Internal server error"
//	@Router			/product/update_stock/{id} [put]
func (p *ProductController) ProductUpdateStock(c *fiber.Ctx) error {
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

	err := p.ProductService.ProductUpdateStock(&stockUpdate)
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
//	@Summary		Update Product
//	@Description	Updates the given product and returns the updated product.
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			product				body		models.ProductUpdate	true	"Product update details"
//	@Success		200					{object}	models.Response			"Product updated successfully"
//	@Failure		400					{object}	models.Response			"Bad Request"
//	@Failure		401					{object}	models.Response			"Auth is required"
//	@Failure		403					{object}	models.Response			"Not Authorized"
//	@Failure		404					{object}	models.Response			"Product not found"
//	@Failure		422					{object}	models.Response			"Model invalid"
//	@Failure		500					{object}	models.Response			"Internal server error"
//	@Router			/product/update [put]
func (p *ProductController) ProductUpdate(c *fiber.Ctx) error {
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

	err := p.ProductService.ProductUpdate(&productUpdate)
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
//	@Summary		Delete Product
//	@Description	Deletes the given product with the given id.
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id					path		string			true	"ID of the product"
//	@Success		200					{object}	models.Response	"Product deleted with success"
//	@Failure		400					{object}	models.Response	"Bad Request"
//	@Failure		401					{object}	models.Response	"Auth is required"
//	@Failure		403					{object}	models.Response	"Not Authorized"
//	@Failure		404					{object}	models.Response	"Product not found"
//	@Failure		500					{object}	models.Response	"Internal server error"
//	@Router			/product/delete/{id} [delete]
func (p *ProductController) ProductDelete(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "ID is required",
		})
	}

	err := p.ProductService.ProductDelete(id)
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
//	@Summary		Create Product
//	@Description	Creates a new product in the specified workplace.
//	@Tags			Product
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			product				body		models.ProductCreate	true	"Details of the product to create"
//	@Success		200					{object}	models.Response			"Product created successfully"
//	@Failure		400					{object}	models.Response			"Bad Request"
//	@Failure		401					{object}	models.Response			"Auth is required"
//	@Failure		403					{object}	models.Response			"Not Authorized"
//	@Failure		422					{object}	models.Response			"Model invalid"
//	@Failure		500					{object}	models.Response			"Internal server error"
//	@Router			/product/create [post]
func (p *ProductController) ProductCreate(c *fiber.Ctx) error {
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

	productCreated, err := p.ProductService.ProductCreate(&productCreate)
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

