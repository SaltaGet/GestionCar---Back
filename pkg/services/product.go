package services

import (
	"errors"

	"github.com/DanielChachagua/GestionCar/pkg/models"
	"github.com/DanielChachagua/GestionCar/pkg/repositories"
	"gorm.io/gorm"
)

func (p *ProductService) ProductGetByID(id string) (*models.Product, error) {
	product, err := p.ProductRepository.ProductGetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "Elemento no encontrado", err)
		}
		return nil, models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return product, nil
}

func (p *ProductService) ProductGetByIdentifier(identifier string) (*[]models.Product, error) {
	product, err := p.ProductRepository.ProductGetByIdentifier(identifier)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, models.ErrorResponse(404, "Elemento no encontrado", err)
		}
		return nil, models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return product, nil
}

func (p *ProductService) ProductGetAll(workplace string) (*[]models.Product, error) {
	products, err := p.ProductRepository.ProductGetAll()
	if err != nil {
		return nil, models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return products, nil
}

func (p *ProductService) ProductGetByName(name string) (*[]models.Product, error) {
	products, err := p.ProductRepository.ProductGetByName(name)
	if err != nil {
		return nil, models.ErrorResponse(500, "Error al obtener productos", err)
	}
	return products, nil
}

func (p *ProductService) ProductCreate(product *models.ProductCreate, workplace string) (string, error) {
	id, err := p.ProductRepository.CreateElement(product, workplace)
	if err != nil {
		return "", models.ErrorResponse(500, "Error al crear producto", err)
	}
	return id, nil
}

func (p *ProductService) ProductUpdate(product *models.ProductUpdate, workplace string) error {
	err := p.ProductRepository.UpdateElement(product, workplace)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Elemento no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al actualizar cliente", err)
	}
	return nil
}

func (p *ProductService) ProductUpdateStock(id string, stock *models.StockUpdate, method string, workplace string) error {
	product, part, err := p.productRepository.ProductGetByID(id, workplace)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Elemento no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al actualizar stock", err)
	}
	switch method {
	case "update":
		if stock.Stock < 0 {
			return models.ErrorResponse(400, "El stock no puede ser negativo", nil)
		}
		return p.ProductRepository.UpdateStock(stock.Stock, id, workplace)
	case "add":
		if stock.Stock <= 0{
			return models.ErrorResponse(400, "El stock debe ser mayor a 0", nil)
		}
		return p.ProductRepository.AddToStock(id, stock.Stock, workplace)
	case "subtract":
		if stock.Stock <= 0{
			return models.ErrorResponse(400, "El stock debe ser mayor a 0", nil)
		}
		if (part != nil && part.Stock < stock.Stock) || (product != nil && product.Stock < stock.Stock) {
			return models.ErrorResponse(400, "El stock no puede ser negativo", nil)
		}
		return p.ProductRepository.SubtractFromStockToStock(id, stock.Stock, workplace)
	
	default:
		return models.ErrorResponse(500, "Método de actualización no soportado", err)
	}
}

func (p *ProductService) ProductDelete(id string, workplace string) error {
	err := repositories.Repo.DeleteElement(id, workplace)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.ErrorResponse(404, "Producto no encontrado", err)
		}
		return models.ErrorResponse(500, "Error al eliminar producto", err)
	}
	return nil
}