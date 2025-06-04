package controllers

import (
	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/gofiber/fiber/v2"
)

func (r *ResumeController) ExpenseResumeCreate(c *fiber.Ctx) error {
	resume := &models.ResumeExpenseCreate{}
	if err := c.BodyParser(resume); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Bad Request",
		})
	}
	if err := resume.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	id, err := r.ResumeService.ExpenseResumeCreate(resume)
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

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Status:  true,
		Body:    id,
		Message: "Resumen creado exitosamente",
	})
}

func (r *ResumeController) ExpenseResumeGetByDateBetween(c *fiber.Ctx) error {
	fromDate := c.Query("fromDate")
	toDate := c.Query("toDate")
	if fromDate == "" || toDate == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "fromDate and toDate are required",
		})
	}

	resumes, err := r.ResumeService.GetExpenseResumeByDateBetween(fromDate, toDate)
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

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Status:  true,
			Body:    resumes,
			Message: "Resumen obtenido exitosamente",
	})
}

func (r *ResumeController) ExpenseResumeGetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "id is required",
		})
	}

	resume, err := r.ResumeService.ExpenseResumeGetByID(id)
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

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Status:  true,
			Body:    resume,
			Message: "Resumen obtenido exitosamente",
	})
}

func (r *ResumeController) ExpenseResumeUpdate(c *fiber.Ctx) error {
	resume := &models.ResumeExpenseUpdate{}
	if err := c.BodyParser(resume); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Bad Request",
		})
	}
	if err := resume.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	err := r.ResumeService.ResumeExpenseUpdate(resume)
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

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Status:  true,
			Body:    nil,
			Message: "Resumen actualizado exitosamente",
	})
}


// INCOME

func (r *ResumeController) IncomeResumeCreate(c *fiber.Ctx) error {
	resume := &models.ResumeIncomeCreate{}
	if err := c.BodyParser(resume); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Bad Request",
		})
	}
	if err := resume.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	id, err := r.ResumeService.CreateIncomeResume(resume)
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

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Status:  true,
		Body:    id,
		Message: "Resumen creado exitosamente",
	})
}

func (r *ResumeController) IncomeResumeGetByDateBetween(c *fiber.Ctx) error {
	fromDate := c.Query("from_date")
	toDate := c.Query("to_date")
	if fromDate == "" || toDate == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "from_date and to_date are required",
		})
	}

	resume, err := r.ResumeService.IncomeResumeGetByDateBetween(fromDate, toDate)
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

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Status:  true,
			Body:    resume,
			Message: "Resumen obtenido exitosamente",
	})
}

func (r *ResumeController) IncomeResumeGetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "id is required",
		})
	}

	resume, err := r.ResumeService.IncomeResumeGetByID(id)
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

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Status:  true,
			Body:    resume,
			Message: "Resumen obtenido exitosamente",
	})
}

func (r *ResumeController) IncomeResumeUpdate(c *fiber.Ctx) error {
	resume := &models.ResumeIncomeUpdate{}
	if err := c.BodyParser(resume); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Bad Request",
		})
	}
	if err := resume.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}

	err := r.ResumeService.UpdateIncomeResume(resume)
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

	return c.Status(fiber.StatusOK).JSON(models.Response{
		Status:  true,
			Body:    nil,
			Message: "Resumen actualizado exitosamente",
	})
}
