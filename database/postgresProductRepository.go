package database

import (
	"database/sql"
	"saas-estoque/entity"
)

type PostgresProductRepository struct {
	db *sql.DB
}

func NewPostgresProductRepository(db *sql.DB) *PostgresProductRepository {
	return &PostgresProductRepository{
		db: db,
	}
}

func (r *PostgresProductRepository) Save(product *entity.Product) error {
	query := `
		INSERT INTO products (
			name,
			description,
			price,
			quantity,
			created_at,
			updated_at,
			category_id
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`

	return r.db.QueryRow(
		query,
		product.Name,
		product.Description,
		product.Price,
		product.Quantity,
		product.CreatedAt,
		product.UpdatedAt,
		product.CategoryID,
	).Scan(&product.ID)
}

func (r *PostgresProductRepository) FindById(id int64) (*entity.Product, error) {
	return nil, nil
}

func (r *PostgresProductRepository) Delete(id int64) error {
	return nil
}

func (r *PostgresProductRepository) FindAll() ([]entity.Product, error) {
	return nil, nil
}

func (r *PostgresProductRepository) Patch(product *entity.Product) error {
	return nil
}
func (r *PostgresProductRepository) Update(product *entity.Product) error {
	return nil
}

//func NewProductMemoryRepository() *ProductMemoryRepository {
//	return &ProductMemoryRepository{products: make(map[int64]entity.Product)}

//Save(product *entity.Product) error
