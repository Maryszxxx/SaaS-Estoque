package usercase

import "saas-estoque/entity"

type UserUseCase interface {
	Create(user entity.User) error
	Update(user entity.User) error
	Delete(id int64) error
}

type LoginUserUseCase interface {
	Login(email, password string) (string, error) //caso retorne um jwt
}
