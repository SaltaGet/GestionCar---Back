package controllers

import (
	"github.com/DanielChachagua/GestionCar/models"
	"github.com/DanielChachagua/GestionCar/services"
	"github.com/gofiber/fiber/v2"
)

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

