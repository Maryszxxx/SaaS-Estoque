package repository

import "saas-estoque/entity"

type ProductRepository interface {
	Save(product *entity.Product) error
	FindByID(id int64) (*entity.Product, error)
	FindAll() ([]entity.Product, error)
	FindDeletedByID(id int64) (*entity.Product, error)
	Update(product *entity.Product) error
	SoftDelete(id int64) error
	Restore(id int64) error
}
type CategoryRepository interface {
	Save(category *entity.Category) error
	FindByID(id int64) (*entity.Category, error)
	FindAll() ([]entity.Category, error)
	FindDeletedByID(id int64) (*entity.Category, error)
	Update(category *entity.Category) error
	SoftDelete(id int64) error
	Restore(id int64) error
}

type StockMovementRepository interface {
	StockIn(movement *entity.StockMovement) error
	StockOut(movement *entity.StockMovement) error
	Adjustment(movement *entity.StockMovement, NewQuantity int64) error
}
