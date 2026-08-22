package usercase

import (
	"errors"
	"saas-estoque/entity"
	"saas-estoque/repository"
	"time"
)

var ErrRestoreWindowExpired = errors.New("the 30-day restore window has expired")

type ProductService struct {
	repository repository.ProductRepository
}
type CategoryService struct {
	repository repository.CategoryRepository
}

func NewCategoryService(repository repository.CategoryRepository) *CategoryService {
	return &CategoryService{
		repository: repository,
	}
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

func (p *ProductService) Patch(id int64, name *string, description *string, price *float64, quantity *int64, categoryID *int64) error {

	product, err := p.repository.FindByID(id)
	if err != nil {
		return err
	}

	if name != nil {
		product.Name = *name
	}

	if description != nil {
		product.Description = *description
	}

	if price != nil {
		product.Price = *price
	}

	if quantity != nil {
		product.Quantity = *quantity
	}

	if categoryID != nil {
		product.CategoryID = *categoryID
	}

	return p.repository.Update(product)
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

func (p *ProductService) Delete(id int64) error {
	return p.repository.SoftDelete(id)
}

func (p *ProductService) Restore(id int64) error {
	product, err := p.repository.FindDeletedByID(id)
	if err != nil {
		return err
	}

	if product.DeletedAt == nil || product.DeletedAt.Before(time.Now().Add(-30*24*time.Hour)) {
		return ErrRestoreWindowExpired
	}

	return p.repository.Restore(id)
}

// category

func (c *CategoryService) Create(name string) (*entity.Category, error) {
	category, err := entity.NewCategoryProduct(name)
	if err != nil {
		return nil, err
	}
	if err := c.repository.Save(category); err != nil {
		return nil, err
	}
	return category, nil
}
func (c *CategoryService) Update(id int64, name *string) error {
	category, err := c.repository.FindByID(id)
	if err != nil {
		return err
	}
	if name != nil {
		category.Name = *name
	}
	return c.repository.Update(category)
}

func (c *CategoryService) FindByID(id int64) (*entity.Category, error) {
	return c.repository.FindByID(id)
}
func (c *CategoryService) FindAll() ([]entity.Category, error) {
	return c.repository.FindAll()
}

func (c *CategoryService) Delete(id int64) error {
	return c.repository.SoftDelete(id)
}

func (c *CategoryService) Restore(id int64) error {
	category, err := c.repository.FindDeletedByID(id)
	if err != nil {
		return err
	}

	if category.DeletedAt == nil || category.DeletedAt.Before(time.Now().Add(-30*24*time.Hour)) {
		return ErrRestoreWindowExpired
	}

	return c.repository.Restore(id)
}

// StockMovementRequest
type StockMovementService struct {
	movementRepository repository.StockMovementRepository
	productRepository  repository.ProductRepository
}

//func NewStockMovementService(movementRepository repository.StockMovementRepository, productRepository repository.ProductRepository) *StockMovementService {
//	return &StockMovementService{
//		movementRepository: movementRepository,
//		productRepository:  productRepository,
//	}
//}

func (s *StockMovementService) StockIn(productID int64, quantity int64, userID int64) error {
	if quantity <= 0 {
		return errors.New("quantity must be greater than zero zero")
	}

	_, err := s.productRepository.FindByID(productID)
	if err != nil {
		return err
	}

	movement := &entity.StockMovement{
		ProductID: productID,
		Quantity:  quantity,
		Type:      "ENTRY",
		UserID:    userID,
		CreatedAt: time.Now(),
	}
	return s.movementRepository.StockIn(movement)

}

func (s *StockMovementService) StockOut(productID int64, quantity int64, userID int64) error {
	product, err := s.productRepository.FindByID(productID)
	if err != nil {
		return err
	}
	if quantity == 0 {
		return errors.New("quantity is zero")
	}
	if product.Quantity < quantity {
		return errors.New("quantity is low")
	}

	movement := &entity.StockMovement{
		ProductID: productID,
		Quantity:  quantity,
		Type:      "EXIT",
		UserID:    userID,
		CreatedAt: time.Now(),
	}

	return s.movementRepository.StockOut(movement)
}

func (s *StockMovementService) Adjustment(productID int64, newQuantity int64, userID int64) error {
	if newQuantity < 0 {
		return errors.New("newQuantity cannot be negative")
	}
	product, err := s.productRepository.FindByID(productID)
	if err != nil {
		return err
	}
	difference := newQuantity - product.Quantity

	if difference == 0 {
		return errors.New("stock already has the quantity")
	}
	movement := &entity.StockMovement{
		ProductID: productID,
		Quantity:  difference,
		Type:      "ADJUSTMENT",
		UserID:    userID,
		CreatedAt: time.Now(),
	}
	return s.movementRepository.Adjustment(movement, newQuantity)
}
