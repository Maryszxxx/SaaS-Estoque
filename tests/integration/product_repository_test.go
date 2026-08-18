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
