package memory

import (
	"errors"
	"saas-estoque/entity"
)

type UserMemoryRepository struct {
	users  map[int64]entity.User
	nextID int64
}

func NewUserMemoryRepository() *UserMemoryRepository {
	return &UserMemoryRepository{users: make(map[int64]entity.User)}
}

func (u *UserMemoryRepository) Save(user *entity.User) error {
	u.nextID++
	user.ID = u.nextID
	u.users[user.ID] = *user
	return nil
}

func (u *UserMemoryRepository) FindById(userID int64) (*entity.User, error) {
	user, ok := u.users[userID]
	if !ok {
		return nil, errors.New("User not found")
	}
	return &user, nil
}
func (u *UserMemoryRepository) FindByEmail(email string) (*entity.User, error) {
	for _, user := range u.users {
		if user.Email == email {
			return &user, nil
		}
	}
	return nil, errors.New("Email not found")
}

func (u *UserMemoryRepository) Update(user *entity.User) error {
	if _, ok := u.users[user.ID]; ok {
		u.users[user.ID] = *user
		return nil
	}
	return errors.New("User not found")

}
func (u *UserMemoryRepository) Delete(id int64) error {
	_, ok := u.users[id]
	if !ok {
		return errors.New("User not found")
	}
	delete(u.users, id)
	return nil
}
