package ports

import "github.com/DanielChachagua/GestionCar/pkg/models"

type ServiceService interface {
	ServiceGetByID(id string) (service *models.Service, err error)
	ServiceGetByName(name string) (service *models.Service, err error)
	ServiceGetAll() (services *[]models.Service, err error)
	ServiceCreate(service *models.ServiceCreate) (id string, err error)
	ServiceUpdate(service *models.ServiceUpdate) (err error)
	ServiceDeleteByID(id string) error
}

type ServiceRepository interface {
	ServiceGetByID(id string) (service *models.Service, err error)
	ServiceGetByName(name string) (exist bool, err error)
	ServiceGetAll() (services *[]models.Service, err error)
	ServiceCreate(service *models.ServiceCreate) (id string, err error)
	ServiceUpdate(service *models.ServiceUpdate) (err error)
	ServiceDeleteByID(id string) error
}
