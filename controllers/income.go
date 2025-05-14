package controllers

import (
	"github.com/DanielChachagua/GestionCar/models"
	"github.com/DanielChachagua/GestionCar/services"
	"github.com/gofiber/fiber/v2"
)

func GetIncomeByID(c *fiber.Ctx) error {
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

	laundry, workshop, err := services.GetIncomeByID(id, workplace.Identifier)
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
			Message: "Ingreso obtenido con éxito",
		})
	}

	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    workshop,
		Message: "Ingreso obtenido con éxito",
	})
}

func GetAllIncomes(c *fiber.Ctx) error {
	workplace := c.Locals("workplace").(*models.Workplace)
	if workplace == nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Workplace is required",
		})
	}

	laundry, workshop, err := services.GetAllIncomes(workplace.Identifier)
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
			Message: "Ingresos obtenidos con éxito",
		})
	}

	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    workshop,
		Message: "Ingresos obtenidos con éxito",
	})
}

func GetIncomeToday(c *fiber.Ctx) error {
	workplace := c.Locals("workplace").(*models.Workplace)
	if workplace == nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Workplace is required",
		})
	}

	laundry, workshop, err := services.GetIncomeToday(workplace.Identifier)
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
			Message: "Ingresos obtenidos con éxito",
		})
	}

	return c.Status(200).JSON(models.Response{
		Status:  true,
		Body:    workshop,
		Message: "Ingresos obtenidos con éxito",
	})
}

func CreateIncome(c *fiber.Ctx) error {
	var incomeCreate models.IncomeCreate
	if err := c.BodyParser(&incomeCreate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Invalid request",
		})
	}
	if err := incomeCreate.Validate(); err != nil {
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

	id, err := services.CreateIncome(&incomeCreate, workplace.Identifier)
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
		Message: "Ingreso creado con éxito",
	})
}

func UpdateIncome(c *fiber.Ctx) error {
	var incomeUpdate models.IncomeUpdate
	if err := c.BodyParser(&incomeUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Invalid request",
		})
	}
	if err := incomeUpdate.Validate(); err != nil {
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

	err := services.UpdateIncome(&incomeUpdate, workplace.Identifier)
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
		Message: "Ingreso editado con éxito",
	})
}

func DeleteIncome(c *fiber.Ctx) error {
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

	err := services.DeleteIncome(id, workplace.Identifier)
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
		Message: "Ingreso eliminado con éxito",
	})
}
