package integration

import (
	"saas-estoque/database"
	"saas-estoque/usercase"
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestUserService_Create(t *testing.T) {
	dsn := "host=localhost port=5433 user=postgres password=postgres dbname=saas_estoque_test sslmode=disable"

	db, err := database.ConnectPostgres(dsn)
	if err != nil {
		t.Fatalf("failed to connect to test database: %v", err)
	}
	t.Cleanup(func() {
		if err := db.Close(); err != nil {
			t.Errorf("failed to close database: %v", err)
		}
	})

	repository := database.NewPostgresUserRepository(db)
	service := usercase.NewUserService(repository)

	user, err := service.Create(
		"Maria",
		"maria@test.com",
		"123456",
		"EMPLOYEE",
	)
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}
	t.Cleanup(func() {
		if err := repository.Delete(user.ID); err != nil {
			t.Errorf("failed to cleanup user: %v", err)
		}
	})
	if user.ID == 0 {
		t.Fatal("expected user ID to be generated")
	}

	if user.Name != "Maria" {
		t.Errorf("expected name %q, got %q", "Maria", user.Name)
	}

	if user.Email != "maria@test.com" {
		t.Errorf("expected email %q, got %q", "maria@test.com", user.Email)
	}

	if user.Role != "EMPLOYEE" {
		t.Errorf("expected role %q, got %q", "EMPLOYEE", user.Role)
	}

	if user.PasswordHash == "123456" {
		t.Fatal("password must not be stored as plain text")
	}
	err = bcrypt.CompareHashAndPassword(
		[]byte(user.PasswordHash),
		[]byte("123456"),
	)

	if err != nil {
		t.Fatal("stored password hash does not match original password")
	}

}

func TestUserService_Login(t *testing.T) {
	dsn := "host=localhost port=5433 user=postgres password=postgres dbname=saas_estoque_test sslmode=disable"

	db, err := database.ConnectPostgres(dsn)
	if err != nil {
		t.Fatalf("failed to connect to test database: %v", err)
	}

	t.Cleanup(func() {
		if err := db.Close(); err != nil {
			t.Errorf("failed to close database: %v", err)
		}
	})

	repository := database.NewPostgresUserRepository(db)
	service := usercase.NewUserService(repository)

	user, err := service.Create(
		"Maria Login",
		"login@test.com",
		"123456",
		"EMPLOYEE",
	)
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	t.Cleanup(func() {
		if err := repository.Delete(user.ID); err != nil {
			t.Errorf("failed to cleanup user: %v", err)
		}
	})

	foundUser, err := service.Login(
		"login@test.com",
		"123456",
	)

	if err != nil {
		t.Fatalf("expected login to succeed: %v", err)
	}

	if foundUser.ID != user.ID {
		t.Errorf(
			"expected user ID %d, got %d",
			user.ID,
			foundUser.ID,
		)
	}

	if foundUser.Email != "login@test.com" {
		t.Errorf(
			"expected email %q, got %q",
			"login@test.com",
			foundUser.Email,
		)
	}
}
