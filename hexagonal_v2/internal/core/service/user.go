package service

import (
	"hexagonal_v2/internal/core/domain"
	"hexagonal_v2/internal/core/port"
)

type userService struct {
	user port.UserService
}

func NewUserService(user port.UserRepository) port.UserService {
	return &userService{user: user}
}

func (u *userService) CreateUser(user *domain.User) error {
	return u.user.CreateUser(user)
}

func (u *userService) GetUserByID(id int64) (*domain.User, error) {
	return u.user.GetUserByID(id)
}

func (u *userService) GetUsers() ([]*domain.User, error) {
	return u.user.GetUsers()
}

func (u *userService) UpdateUser(user *domain.User) error {
	return u.user.UpdateUser(user)
}

func (u *userService) DeleteUser(id int64) error {
	return u.user.DeleteUser(id)
}
