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

func (p *ProductService) Update(id int64, name string, description string, price float64, quantity int64, categoryID int64) error {
	product := entity.Product{
		ID:          id,
		Name:        name,
		Description: description,
		Price:       price,
		Quantity:    quantity,
		CategoryID:  categoryID,
	}
	return p.repository.Update(&product)
}

func (p *ProductService) Patch(id int64, name string, description string, price float64, quantity int64, categoryID int64) error {
	product := entity.Product{
		ID:          id,
		Name:        name,
		Description: description,
		Price:       price,
		Quantity:    quantity,
		CategoryID:  categoryID,
	}
	return p.repository.Patch(&product)
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

	return p.repository.Save(product)
}

func (p *ProductService) FindByID(id int64) (*entity.Product, error) {
	return p.repository.FindByID(id)
}

func (p *ProductService) FindAll() ([]entity.Product, error) {
	return p.repository.FindAll()
}
