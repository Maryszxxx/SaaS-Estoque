package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"saas-estoque/entity"
)

type PostgresProductRepository struct {
	db *sql.DB
}

var (
	DB_HOST     = os.Getenv("DB_HOST")
	DB_USERNAME = os.Getenv("DB_USERNAME")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_DATABASE = os.Getenv("DB_DATABASE")
	DB_PORT     = os.Getenv("DB_PORT")
	DBNAME      = os.Getenv("DB_NAME")
	SSLMODE     = os.Getenv("SSLMODE")
)

func ConnectPostgresProductRepository() *sql.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=%s", DB_HOST, DB_USERNAME, DB_PASSWORD, DB_DATABASE, DB_PORT,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	return db
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

func (r *PostgresProductRepository) FindByID(id int64) (*entity.Product, error) {
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
