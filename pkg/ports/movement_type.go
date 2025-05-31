package ports

import "github.com/DanielChachagua/GestionCar/pkg/models"

type MovementTypeService interface {
	GetMovementTypeByID(id string) (movementType *models.MovementType, err error)
	GetAllMovementTypes(isIncome bool) (movementTypes *[]models.MovementType, err error)
	CreateMovementType(movementType *models.MovementTypeCreate) (id string, err error)
	UpdateMovementType(movementTypeUpdate *models.MovementTypeUpdate) (err error)
	DeleteMovementType(id string) (err error)
}

type MovementTypeRepository interface {
	GetMovementTypeByID(id string) (movementType *models.MovementType, err error)
	GetAllMovementTypes(isIncome bool) (movementTypes *[]models.MovementType, err error)
	CreateMovementType(movementType *models.MovementTypeCreate) (id string, err error)
	UpdateMovementType(movementTypeUpdate *models.MovementTypeUpdate) (err error)
	DeleteMovementType(id string) (err error)
}
