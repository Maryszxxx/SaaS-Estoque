package integration

import (
	"database/sql"
	"saas-estoque/database"
	"saas-estoque/entity"
	"testing"
	"time"
)

func TestUserRepository(t *testing.T) {
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

	repository := database.NewPostgresUserRepository(db)

	//CREATE
	user := &entity.User{
		Name:         "Usuário integração",
		Email:        "integracao@test.com",
		Role:         entity.RoleEmployee,
		PasswordHash: "hash-fake-para-teste",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	err = repository.Save(user)
	if err != nil {
		t.Fatalf("failed to save user: %v", err)
	}
	if user.ID == 0 {
		t.Fatalf("failed to save user")
	}

	//READ BY ID

	foundUser, err := repository.FindByID(user.ID)
	if err != nil {
		t.Fatalf("failed to find user by id: %v", err)
	}
	if foundUser.Name != user.Name {
		t.Errorf(
			"expected name %q, got %q",
			user.Name,
			foundUser.Name,
		)
	}
	if foundUser.Email != user.Email {
		t.Errorf(
			"expected email %q, got %q",
			user.Email,
			foundUser.Email,
		)
	}
	if foundUser.Role != user.Role {
		t.Errorf(
			"expected role %q, got %q",
			user.Role,
			foundUser.Role,
		)
	}

	// READ BY EMAIL
	foundByEmail, err := repository.FindByEmail(user.Email)
	if err != nil {
		t.Fatalf("failed to find user by email: %v", err)
	}

	if foundByEmail.ID != user.ID {
		t.Errorf(
			"expected user ID %d, got %d",
			user.ID,
			foundByEmail.ID,
		)
	}

	// UPDATE
	user.Name = "Usuário atualizado"
	user.Email = "atualizado@test.com"
	user.Role = entity.RoleAdmin
	user.PasswordHash = "novo-hash"

	err = repository.Update(user)
	if err != nil {
		t.Fatalf("failed to update user: %v", err)
	}

	updatedUser, err := repository.FindByID(user.ID)
	if err != nil {
		t.Fatalf("failed to find updated user: %v", err)
	}

	if updatedUser.Name != "Usuário atualizado" {
		t.Errorf(
			"expected updated name %q, got %q",
			"Usuário atualizado",
			updatedUser.Name,
		)
	}

	if updatedUser.Email != "atualizado@test.com" {
		t.Errorf(
			"expected updated email %q, got %q",
			"atualizado@test.com",
			updatedUser.Email,
		)
	}

	if updatedUser.Role != entity.RoleAdmin {
		t.Errorf(
			"expected role %q, got %q",
			entity.RoleAdmin,
			updatedUser.Role,
		)
	}

	if updatedUser.PasswordHash != "novo-hash" {
		t.Errorf(
			"expected password hash %q, got %q",
			"novo-hash",
			updatedUser.PasswordHash,
		)
	}

	// DELETE
	err = repository.Delete(user.ID)
	if err != nil {
		t.Fatalf("failed to delete user: %v", err)
	}

	// VERIFY DELETE
	_, err = repository.FindByID(user.ID)
	if err == nil {
		t.Fatal("expected error when finding deleted user")
	}

}
