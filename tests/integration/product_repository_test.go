package integration

import (
	"fmt"
	"os"
	"testing"
	"time"

	"saas-estoque/database"
	"saas-estoque/entity"

	_ "github.com/lib/pq"
)

func TestProductRepository(t *testing.T) {
	password := os.Getenv("PASSWORD")

	dsn := fmt.Sprintf(
		"host=localhost port=5432 user=postgres password=%s dbname=Saas-Estoque-Test sslmode=disable",
		password,
	)

	db, err := database.ConnectPostgres(dsn)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	repository := database.NewPostgresProductRepository(db)

	//create
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
		t.Fatal(err)
	}
	if product.ID == 0 {
		t.Error("expected product ID to be generated")
	}

	//read by ID
	foundProduct, err := repository.FindByID(product.ID)
	if err != nil {
		t.Fatal(err)
	}
	if foundProduct.Name != product.Name {
		t.Errorf("expected %s got %s", product.Name, foundProduct.Name)
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

func TestProductRepository_SaveAndFindByID(t *testing.T) {
	dsn := fmt.Sprintf(
		"host=localhost port=5432 user=postgres password=%s dbname=%s sslmode=disable",
		os.Getenv("TEST_DB_PASSWORD"),
		"Saas-Estoque-Test",
	)

	db, err := database.ConnectPostgres(dsn)
	if err != nil {
		t.Fatalf("failed to connect to test database: %v", err)
	}

	defer db.Close()

	repository := database.NewPostgresProductRepository(db)

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
}
