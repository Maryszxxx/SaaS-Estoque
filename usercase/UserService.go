package usercase

import (
	"errors"
	"saas-estoque/entity"
	"saas-estoque/repository"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *UserService {
	return &UserService{
		userRepository: repository,
	}
}

// implementando user

func (p *UserService) Create(name, email, password, role string) error {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user, err := entity.NewUser(
		name,
		email,
		string(hash),
		role,
	)
	if err != nil {
		return err
	}

	return p.userRepository.Save(user)
}

func (u *UserService) FindByEmail(email string) (*entity.User, error) {
	return u.userRepository.FindByEmail(email)
}

func (u *UserService) FindById(ID int64) (*entity.User, error) {
	return u.userRepository.FindByID(ID)
}

func (u *UserService) Delete(ID int64) error {
	return u.userRepository.Delete(ID)

}

func (u *UserService) Update(ID int64, name, email, password, role string) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := &entity.User{
		ID:           ID,
		Name:         name,
		Email:        email,
		Role:         role,
		PasswordHash: string(hash),
	}
	return u.userRepository.Save(user)
}

// patch sem senha
func (u *UserService) Patch(ID int64, name *string, email *string, role *string) error {

	user, err := u.userRepository.FindByID(ID)
	if err != nil {
		return err
	}
	if name != nil {
		user.Name = *name
	}
	if email != nil {
		user.Email = *email
	}

	if role != nil {
		switch *role {
		case entity.RoleAdmin:
			user.Role = *role
		default:
			return errors.New("invalid role")
		}
	}
	return u.userRepository.Save(user)
}

// patch apenas pra senha
func (u *UserService) ChangePassword(ID int64, newPassword, oldPassword *string) error {

	user, err := u.userRepository.FindByID(ID)

	if err != nil {
		return err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(*oldPassword)); err != nil {
		return errors.New("Senha atual incorreta")

	}

	newHash, err := bcrypt.GenerateFromPassword(
		[]byte(*newPassword), bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}
	user.PasswordHash = string(newHash)
	return u.userRepository.Update(user)
}
