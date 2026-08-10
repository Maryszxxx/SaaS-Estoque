package usercase

import (
	"saas-estoque/entity"
	"saas-estoque/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

// implementando user

func (p *UserService) CreateUser(name, email, passwordHash, role string) error {

	user, err := entity.NewUser(
		name,
		email,
		passwordHash,
		role,
	)
	if err != nil {
		return err
	}

	return p.userRepository.Save(user)
}
