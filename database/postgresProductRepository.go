package database

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"saas-estoque/entity"
	"time"

	"github.com/joho/godotenv"
)

type PatchProductRequest struct {
	Name        *string
	Description *string
	Price       *float64
	Quantity    *int64
	CategoryID  *int64
}
type PostgresProductRepository struct {
	db *sql.DB
}

func ConnectPostgres(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func ConnectPostgresRepository() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}

	DB_HOST := os.Getenv("DB_HOST")
	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_DATABASE := os.Getenv("DB_DATABASE")
	DB_PORT := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		DB_HOST,
		DB_PORT,
		DB_USERNAME,
		DB_PASSWORD,
		DB_DATABASE,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("erro ao abrir conexão com banco de dados", err)
	}

	for i := 0; i < 10; i++ {
		err := db.Ping()

		if err == nil {
			log.Println("Connected to database")
			return db
		}
		log.Println("aguardando postgresql")
		time.Sleep(2 * time.Second)
	}
	log.Fatal("erro ao conectar com banco de dados", err)
	return nil
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
	query := `SELECT id, description, name, price, quantity, created_at, updated_at, category_id FROM products WHERE id = $1`
	product := &entity.Product{}

	err := r.db.QueryRow(query, id).Scan(&product.ID, &product.Description, &product.Name, &product.Price, &product.Quantity, &product.CreatedAt, &product.UpdatedAt, &product.CategoryID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("no rows found")
		}
		return nil, err
	}
	return product, nil

}

func (r *PostgresProductRepository) Delete(id int64) error {
	query := `DELETE FROM products WHERE id = $1`
	results, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := results.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("no product found with that id")
	}

	return nil
}

func (r *PostgresProductRepository) FindAll() ([]entity.Product, error) {
	query := `SELECT id, description, name, price, quantity, created_at, updated_at, category_id FROM products`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var products []entity.Product

	for rows.Next() {
		product := entity.Product{}

		err := rows.Scan(&product.ID, &product.Description, &product.Name, &product.Price, &product.Quantity, &product.CreatedAt, &product.UpdatedAt, &product.CategoryID)

		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil

}

func (r *PostgresProductRepository) Update(product *entity.Product) error {
	query := `UPDATE products SET name = $1, description = $2, price = $3, quantity = $4, category_id = $5, updated_at = $6 WHERE id = $7 `
	result, err := r.db.Exec(query, product.Name, product.Description, product.Price, product.Quantity, product.CategoryID, time.Now(), product.ID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("product no found")
	}
	return nil
}
