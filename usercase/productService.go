package usercase

import (
	"saas-estoque/entity"
	"saas-estoque/repository"
)

type ProductService struct {
	repository repository.ProductRepository
}

func NewProductService(repository repository.ProductRepository) *ProductService {
	return &ProductService{
		repository: repository,
	}
}

func (p *ProductService) Create(name string, description string, price float64, quantity int64, categoryID int64,
) error {
	product, err := entity.NewProduct(
		name,
		description,
		price,
		quantity,
		categoryID,
	)
	if err != nil {
		return err
	}

	return p.repository.Save(*product)
}

func (p *ProductService) FindById(id int64) (*entity.Product, error) {
	product, err := entity.User{ID: id}
	if err != nil {
		return nil, err
	}
}
