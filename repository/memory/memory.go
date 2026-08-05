package memory

import (
	"errors"
	"saas-estoque/entity"
)

type ProductMemoryRepository struct {
	products map[int64]entity.Product
	nextID   int64
}

func NewProductMemoryRepository() *ProductMemoryRepository {
	return &ProductMemoryRepository{products: make(map[int64]entity.Product)}
}

func (p *ProductMemoryRepository) FindByID(productID int64) (entity.Product, error) {
	product, ok := p.products[productID]
	if !ok {
		return entity.Product{}, errors.New("product not found")
	}
	return product, nil

}

func (p *ProductMemoryRepository) FindAll() ([]entity.Product, error) {
	products := []entity.Product{}
	for _, product := range p.products {
		products = append(products, product)
	}
	return products, nil
}

func (p *ProductMemoryRepository) Delete(productID int64) error {
	_, ok := p.products[productID]
	if ok {
		delete(p.products, productID)
	} else {
		return errors.New("product not found")
	}
	return nil
}

func (p *ProductMemoryRepository) Save(product *entity.Product) error {
	p.nextID++
	product.ID = p.nextID
	p.products[product.ID] = *product
	return nil
}

func (p *ProductMemoryRepository) Update(product *entity.Product) error {
	if _, ok := p.products[product.ID]; !ok {
		return errors.New("product not found")
	}

	p.products[product.ID] = *product

	return nil
}
