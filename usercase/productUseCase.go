package usercase

import (
	"saas-estoque/entity"
	"saas-estoque/repository"
)

type ProductUseCase interface {
	Create(product entity.Product) error
	Update(product entity.Product) error
	Delete(id int64) error
	FindByID(id int64) (*entity.Product, error)
	FindAll() ([]entity.Product, error)
}

type UserUseCase interface {
	Create(user entity.User) error
	Update(user entity.User) error
	Delete(id int64) error
}

type LoginUserUseCase interface {
	Login(email, password string) (string, error) //caso retorne um jwt
}

type ProductService struct {
	repository repository.ProductRepository
}
