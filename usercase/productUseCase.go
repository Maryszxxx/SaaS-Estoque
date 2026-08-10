package usercase

import (
	"saas-estoque/entity"
)

type ProductUseCase interface {
	Create(
		name string,
		description string,
		price float64,
		quantity int64,
		categoryID int64,
	) error
	Update(product entity.Product) error
	Delete(id int64) error
	FindByID(id int64) (*entity.Product, error)
	FindAll() ([]entity.Product, error)
}
