package integration

import (
	"database/sql"
	"testing"
	"time"

	"saas-estoque/database"
	"saas-estoque/entity"

	_ "github.com/lib/pq"
)

func TestProductRepository(t *testing.T) {
	dsn := "host=localhost port=5433 user=postgres password=postgres dbname=saas_estoque_test sslmode=disable"

	db, err := database.ConnectPostgres(dsn)
	if err != nil {
		t.Fatalf("failed to connect to test database: %v", err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			panic(err)
		}
	}(db)

	repository := database.NewPostgresProductRepository(db)

	// CREATE
	product := &entity.Product{
		Name:        "Produto integração",
		Description: "Produto criado pelo teste",
		Price:       100.50,
		Quantity:    10,
		CategoryID:  1,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = repository.Save(product)
	if err != nil {
		t.Fatalf("failed to save product: %v", err)
	}

	if product.ID == 0 {
		t.Fatal("expected product ID to be generated")
	}

	// READ BY ID
	foundProduct, err := repository.FindByID(product.ID)
	if err != nil {
		t.Fatalf("failed to find product: %v", err)
	}

	if foundProduct.Name != product.Name {
		t.Errorf(
			"expected name %q, got %q",
			product.Name,
			foundProduct.Name,
		)
	}

	// READ ALL
	products, err := repository.FindAll()
	if err != nil {
		t.Fatalf("failed to find all products: %v", err)
	}

	if len(products) == 0 {
		t.Fatal("expected at least one product")
	}

	// UPDATE
	product.Name = "Produto atualizado"
	product.Price = 200.00
	product.Quantity = 20

	err = repository.Update(product)
	if err != nil {
		t.Fatalf("failed to update product: %v", err)
	}

	updatedProduct, err := repository.FindByID(product.ID)
	if err != nil {
		t.Fatalf("failed to find updated product: %v", err)
	}

	if updatedProduct.Name != "Produto atualizado" {
		t.Errorf(
			"expected updated name %q, got %q",
			"Produto atualizado",
			updatedProduct.Name,
		)
	}

	if updatedProduct.Price != 200.00 {
		t.Errorf(
			"expected price %.2f, got %.2f",
			200.00,
			updatedProduct.Price,
		)
	}

	if updatedProduct.Quantity != 20 {
		t.Errorf(
			"expected quantity %d, got %d",
			20,
			updatedProduct.Quantity,
		)
	}

	// DELETE
	err = repository.Delete(product.ID)
	if err != nil {
		t.Fatalf("failed to delete product: %v", err)
	}

	// VERIFY DELETE
	_, err = repository.FindByID(product.ID)
	if err == nil {
		t.Fatal("expected error when finding deleted product")
	}
}
