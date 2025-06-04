package controllers

import (
	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/DanielChachagua/GestionCar/pkg/services"
	"github.com/gofiber/fiber/v2"
)


// ClientGetByID godoc
//	@Summary		Get client by id
//	@Description	Get client by id
//	@Tags			Client
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path		string	true	"Id del cliente"
//	@Success		200	{object}	models.Response{body=models.Client}
//	@Failure		400	{object}	models.Response
//	@Failure		401	{object}	models.Response
//	@Failure		403	{object}	models.Response
//	@Failure		404	{object}	models.Response
//	@Failure		500	{object}	models.Response
//	@Router			/client/{id} [get]
func (cl *ClientController) ClientGetByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "ID is required",
		})
	}

	client, err := cl.ClientService.ClientGetByID(id)
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
//	@Summary		Get All Clients
//	@Description	Get All Clients
//	@Tags			Client
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Success		200	{object}	models.Response{body=[]models.Client}
//	@Failure		400	{object}	models.Response
//	@Failure		401	{object}	models.Response
//	@Failure		403	{object}	models.Response
//	@Failure		404	{object}	models.Response
//	@Failure		500	{object}	models.Response
//	@Router			/client/get_all [get]
func (cl *ClientController) ClientGetAll(c *fiber.Ctx) error {
	clients, err := cl.ClientService.ClientGetAll()
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
//	@Summary		Get Client By Name
//	@Description	Get Client By Name
//	@Tags			Client
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			name	query		string	true	"Name"
//	@Success		200		{object}	models.Response{body=[]models.Client}
//	@Failure		400		{object}	models.Response
//	@Failure		401		{object}	models.Response
//	@Failure		403		{object}	models.Response
//	@Failure		404		{object}	models.Response
//	@Failure		500		{object}	models.Response
//	@Router			/client/get_by_name [get]
func (cl *ClientController) ClientGetByName(c *fiber.Ctx) error {
	name := c.Query("name")
	if name == "" || len(name) < 3 {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "El valor no debe de ser vacio o menor a 3 caracteres",
		})
	}

	clients, err := cl.ClientService.ClientGetByName(name)
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
//	@Summary		Actualizar un cliente
//	@Description	Actualizar un cliente
//	@Tags			Client
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			ClientUpdate	body		models.ClientUpdate	true	"Cliente a actualizar"
//	@Success		200				{object}	models.Response
//	@Failure		400				{object}	models.Response
//	@Failure		401				{object}	models.Response
//	@Failure		403				{object}	models.Response
//	@Failure		404				{object}	models.Response
//	@Failure		422				{object}	models.Response
//	@Failure		500				{object}	models.Response
//	@Router			/client/update [put]
func (cl *ClientController) ClientUpdate(c *fiber.Ctx) error {
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
	clientCreated, err := cl.ClientService.ClientUpdate(&clientUpdate)
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

// ClientDelete godoc
//	@Summary		Delete client by ID
//	@Description	Delete client by ID
//	@Tags			Client
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			id	path		string	true	"Id del cliente"
//	@Success		200	{object}	models.Response{body=models.Client}
//	@Failure		400	{object}	models.Response
//	@Failure		401	{object}	models.Response
//	@Failure		403	{object}	models.Response
//	@Failure		404	{object}	models.Response
//	@Failure		500	{object}	models.Response
//	@Router			/client/delete/{id} [delete]
func (cl *ClientController) ClientDelete(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return c.Status(fiber.StatusBadRequest).JSON(models.Response{
			Status:  false,
			Body:    nil,
			Message: "ID is required",
		})
	}

	client, err := cl.ClientService.ClientDelete(id)
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

// CreateClient godoc
//	@Summary		Create client
//	@Description	Create client
//	@Tags			Client
//	@Accept			json
//	@Produce		json
//	@Security		BearerAuth
//	@Param			clientCreate	body		models.ClientCreate	true	"Información del cliente"
//	@Success		200				{object}	models.Response
//	@Failure		400				{object}	models.Response
//	@Failure		401				{object}	models.Response
//	@Failure		403				{object}	models.Response
//	@Failure		422				{object}	models.Response
//	@Failure		500				{object}	models.Response
//	@Router			/client/create [post]
func (cl *ClientController) CreateClient(c *fiber.Ctx) error {
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
	clientCreated, err := cl.ClientService.ClientCreate(&clientCreate)
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
