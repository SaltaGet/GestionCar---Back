package controllers

import (
	"github.com/DanielChachagua/GestionCar/models"
	"github.com/DanielChachagua/GestionCar/services"
	"github.com/gofiber/fiber/v2"
)


// ClientGetByID obtener un cliente por id
// @Summary Obtiene un cliente por su id
// @Description Obtiene un cliente por su id
// @Tags Clientes
// @Accept  json
// @Produce  json
// @Param   id     path    string     true  "Id del cliente"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Failure 404 {object} models.Response
// @Failure 500 {object} models.Response
// @Router /api/v1/clients/{id} [get]
func ClientGetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "ID is required",
		})
	}

	client, err := services.ClientGetByID(id)
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
		Body:    client,
		Message: "Cliente obtenido con éxito",
	})
}

// ClientGetAll godoc
// @Summary     Get All Clients
// @Description Get All Clients
// @Tags        clients
// @Accept      json
// @Produce     json
// @Success     200  {object}  models.Response{body=[]models.Client}
// @Failure     400  {object}  models.Response
// @Failure     404  {object}  models.Response
// @Failure     500  {object}  models.Response
// @Router      /clients [get]
// @Security    BearerAuth
func ClientGetAll(c *fiber.Ctx) error {
	clients, err := services.ClientGetAll()
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
		Body:    clients,
		Message: "Clientes obtenidos con éxito",
	})
}

// ClientGetByName godoc
// @Summary     Get Client By Name
// @Description Get Client By Name
// @Tags        clients
// @Accept      json
// @Produce     json
// @Param       name  query     string  true  "Name"
// @Success     200  {object}  models.Response{body=[]models.Client}
// @Failure     400  {object}  models.Response
// @Failure     404  {object}  models.Response
// @Failure     500  {object}  models.Response
// @Router      /clients/by-name [get]
// @Security    BearerAuth
func ClientGetByName(c *fiber.Ctx) error {
	name := c.Query("name")
	if name == "" || len(name) < 3 {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "El valor no debe de ser vacio o menor a 3 caracteres",
		})
	}

	clients, err := services.ClientGetByName(name)
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
		Body:    clients,
		Message: "Clientes obtenidos con éxito",
	})
}

// ClientUpdate actualiza un cliente
// @Summary     Actualizar un cliente
// @Description Actualizar un cliente
// @Tags        clients
// @Accept      json
// @Produce     json
// @Param       ClientUpdate  body      models.ClientUpdate  true  "Cliente a actualizar"
// @Success     200  {object}  models.Response{body=models.Client}
// @Failure     400  {object}  models.Response
// @Failure     404  {object}  models.Response
// @Failure     500  {object}  models.Response
// @Router      /clients [put]
// @Security    BearerAuth
func ClientUpdate(c *fiber.Ctx) error {
	var clientUpdate models.ClientUpdate
	if err := c.BodyParser(&clientUpdate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Invalid request",
		})
	}
	if err := clientUpdate.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}
	clientCreated, err := services.ClientUpdate(&clientUpdate)
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
		Body:    clientCreated,
		Message: "Cliente actualizado con éxito",
	})
}

// ClientDelete elimina un cliente por su id
// @Summary     Eliminar un cliente
// @Description Eliminar un cliente
// @Tags        clients
// @Accept      json
// @Produce     json
// @Param       id   path      string     true  "Id del cliente"
// @Success     200  {object}  models.Response{body=models.Client}
// @Failure     400  {object}  models.Response
// @Failure     404  {object}  models.Response
// @Failure     500  {object}  models.Response
// @Router      /clients/{id} [delete]
// @Security    BearerAuth
func ClientDelete(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "ID is required",
		})
	}

	client, err := services.ClientDelete(id)
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
		Body:    client,
		Message: "Cliente eliminado con éxito",
	})
}

// CreateClient crea un cliente
// @Summary     Crear un cliente
// @Description Crear un cliente
// @Tags        clients
// @Accept      json
// @Produce     json
// @Param       clientCreate  body      models.ClientCreate  true  "Información del cliente"
// @Success     200           {object}  models.Response{body=models.Client}
// @Failure     400           {object}  models.Response
// @Failure     500           {object}  models.Response
// @Router      /clients [post]
// @Security    BearerAuth
func CreateClient(c *fiber.Ctx) error {
	var clientCreate models.ClientCreate
	if err := c.BodyParser(&clientCreate); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "Invalid request",
		})
	}
	if err := clientCreate.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: err.Error(),
		})
	}
	clientCreated, err := services.ClientCreate(&clientCreate)
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
		Body:    clientCreated,
		Message: "Cliente creado con éxito",
	})
}
