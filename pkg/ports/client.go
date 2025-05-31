package ports

import "github.com/DanielChachagua/GestionCar/pkg/models"

type ClientService interface {
	GetClientByID(id string) (client *models.Client, err error)
	GetClientByName(name string) (clients *[]models.Client, err error)
	GetAllClients() (clients *[]models.Client, err error)
	CreateClient(clientCreate *models.ClientCreate) (id string, err error)
	UpdateClient(clienUpdate *models.ClientUpdate) (err error)
	DeleteClient(id string) (err error)
}

type ClientRepository interface {
	GetClientByID(id string) (client *models.Client, err error)
	GetClientByName(name string) (clients *[]models.Client, err error)
	GetAllClients() (clients *[]models.Client, err error)
	CreateClient(clientCreate *models.ClientCreate) (id string, err error)
	UpdateClient(clienUpdate *models.ClientUpdate) (err error)
	DeleteClient(id string) (err error)
}
