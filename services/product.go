package services

import (
	"errors"

	"github.com/DanielChachagua/GestionCar/models"
	"github.com/DanielChachagua/GestionCar/repositories"
	"gorm.io/gorm"
)

func ProductGetByID(id string, workplace string) (*models.ProductLaundry, *models.PartWorkshop, error) {
	product, part, err := repositories.Repo.GetElementByID(id, workplace)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, models.ErrorResponse(404, "Empleado no encontrado", err)
		}
		return nil, nil, models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return product, part, nil
}

func ProductGetByIdentifier(identifier string, workplace string) (*[]models.ProductLaundry, *[]models.PartWorkshop, error) {
	product, part, err := repositories.Repo.GetElementsByIdentifier(identifier, workplace)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, models.ErrorResponse(404, "Empleado no encontrado", err)
		}
		return nil, nil, models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return product, part, nil
}

func ProductGetAll(workplace string) (*[]models.ProductLaundry, *[]models.PartWorkshop, error) {
	product, part, err := repositories.Repo.GetAllElements(workplace)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, models.ErrorResponse(404, "Empleado no encontrado", err)
		}
		return nil, nil, models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return product, part, nil
}

func ProductGetByName(name string, workplace string) (*[]models.ProductLaundry, *[]models.PartWorkshop, error) {
	product, part, err := repositories.Repo.GetAllElementsByName(name, workplace)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil, models.ErrorResponse(404, "Empleado no encontrado", err)
		}
		return nil, nil, models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return product, part, nil
}

func ProductCreate(product *models.ProductCreate, workplace string) (string, error) {
	id, err := repositories.Repo.CreateElement(product, workplace)
	if err != nil {
		return "", models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return id, nil
}

func ProductUpdate(product *models.ProductUpdate, workplace string) error {
	err := repositories.Repo.UpdateElement(product, workplace)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Empleado no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return nil
}

func ProductUpdateStock(id string, stock *models.StockUpdate, method string, workplace string) error {
	product, part, err := repositories.Repo.GetElementByID(id, workplace)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Elemento no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	switch method {
	case "update":
		if stock.Stock < 0 {
			return models.ErrorResponse(400, "El stock no puede ser negativo", nil)
		}
		return repositories.Repo.UpdateStock(stock.Stock, id, workplace)
	case "add":
		if stock.Stock <= 0{
			return models.ErrorResponse(400, "El stock debe ser mayor a 0", nil)
		}
		return repositories.Repo.AddToStock(id, stock.Stock, workplace)
	case "subtract":
		if stock.Stock <= 0{
			return models.ErrorResponse(400, "El stock debe ser mayor a 0", nil)
		}
		if (part != nil && part.Stock < stock.Stock) || (product != nil && product.Stock < stock.Stock) {
			return models.ErrorResponse(400, "El stock no puede ser negativo", nil)
		}
		return repositories.Repo.SubtractFromStockToStock(id, stock.Stock, workplace)
	
	default:
		return models.ErrorResponse(500, "Método de actualización no soportado", err)
	}
}

func ProductDelete(id string, workplace string) error {
	err := repositories.Repo.DeleteElement(id, workplace)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Empleado no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return nil
}