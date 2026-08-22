package memory

import (
	"errors"
	"saas-estoque/entity"
	"time"
)

type ProductMemoryRepository struct {
	products map[int64]entity.Product
	nextID   int64
}

func (p *ProductMemoryRepository) Patch(product *entity.Product) error {
	if _, ok := p.products[product.ID]; !ok {
		return errors.New("product not found")
	}

	p.products[product.ID] = *product

	return nil
}

func (p *ProductMemoryRepository) FindByID(productID int64) (*entity.Product, error) {
	product, ok := p.products[productID]
	if !ok || product.DeletedAt != nil {
		return nil, errors.New("product not found")
	}
	return &product, nil

}

func (p *ProductMemoryRepository) FindAll() ([]entity.Product, error) {
	products := []entity.Product{}
	for _, product := range p.products {
		if product.DeletedAt == nil {
			products = append(products, product)
		}
	}
	return products, nil
}

func (p *ProductMemoryRepository) SoftDelete(productID int64) error {
	product, ok := p.products[productID]
	if !ok || product.DeletedAt != nil {
		return errors.New("product not found")
	}
	now := time.Now()
	product.DeletedAt = &now
	p.products[productID] = product
	return nil
}

func (p *ProductMemoryRepository) Save(product *entity.Product) error {
	p.nextID++
	product.ID = p.nextID
	p.products[product.ID] = *product
	return nil
}

func (p *ProductMemoryRepository) Update(product *entity.Product) error {
	stored, ok := p.products[product.ID]
	if !ok || stored.DeletedAt != nil {
		return errors.New("product not found")
	}

	p.products[product.ID] = *product

	return nil
}

func (p *ProductMemoryRepository) FindDeletedByID(productID int64) (*entity.Product, error) {
	product, ok := p.products[productID]
	if !ok || product.DeletedAt == nil {
		return nil, errors.New("deleted product not found")
	}
	return &product, nil
}

func (p *ProductMemoryRepository) Restore(productID int64) error {
	product, ok := p.products[productID]
	if !ok || product.DeletedAt == nil {
		return errors.New("deleted product not found")
	}
	product.DeletedAt = nil
	p.products[productID] = product
	return nil
}
