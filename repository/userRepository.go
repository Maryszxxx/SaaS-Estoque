package repository

import "saas-estoque/entity"

type UserRepository interface {
	Save(user *entity.User) error
	FindByID(id int64) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	Update(product *entity.User) error
	Delete(id int64) error
}
