package repository

import "saas-estoque/entity"

type ProductRepository interface {
	Save(product entity.Product) error
	FindByID(id int64) (*entity.Product, error)
	FindAll() ([]entity.Product, error)
	Update(product entity.Product) error
	Delete(id int64) error
}
