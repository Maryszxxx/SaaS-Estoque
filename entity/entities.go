package entity

import (
	"errors"
	"net/mail"
	"time"
)

const (
	RoleAdmin    = "ADMIN"
	RoleEmployee = "EMPLOYEE"
)

type User struct {
	ID           int64
	Name         string
	Email        string
	PasswordHash string
	Role         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Product struct {
	ID          int64
	Description string
	Name        string
	Price       float64
	Quantity    int64
	CreatedAt   time.Time
	UpdatedAt   time.Time
	CategoryID  int64
}

type Category struct {
	ID        int64
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewProduct(name, description string, price float64, quantity, categoryID int64) (*Product, error) {
	now := time.Now()
	product := &Product{
		Name:        name,
		Description: description,
		Price:       price,
		Quantity:    quantity,
		CreatedAt:   now,
		UpdatedAt:   now,
		CategoryID:  categoryID,
	}
	if name == "" {
		return nil, errors.New("name is required")
	}
	if description == "" {
		return nil, errors.New("description is required")
	}
	if price <= 0 {
		return nil, errors.New("the price must be greater than zero")
	}
	if quantity <= 0 {
		return nil, errors.New("the quantity can't be negative")
	}
	if categoryID <= 0 {
		return nil, errors.New("CategoryID must be greater than zero")
	}
	return product, nil
}

func NewUser(name, email, passwordHash, role string) (*User, error) {
	now := time.Now()
	user := &User{
		Name:         name,
		Email:        email,
		Role:         role,
		PasswordHash: passwordHash,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
	if name == "" {
		return nil, errors.New("name is required")
	}
	_, err := mail.ParseAddress(email)
	if err != nil {
		return nil, errors.New("email is required")
	}

	if passwordHash == "" {
		return nil, errors.New("password is required")
	}

	switch role {
	case RoleAdmin, RoleEmployee:
	default:
		return nil, errors.New("role is required")
	}
	return user, nil
}
