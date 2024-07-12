package service

import (
	"hexagonal_v2/internal/core/domain"
	"hexagonal_v2/internal/core/port"
)

type UserService struct {
	user port.UserService
}

func NewUserService(user port.UserService) *UserService {
	return &UserService{user: user}
}

func (u *UserService) CreateUser(user *domain.User) error {
	return u.user.CreateUser(user)
}

func (u *UserService) GetUserByID(id int) (*domain.User, error) {
	return u.user.GetUserByID(id)
}

func (u *UserService) GetUsers() ([]*domain.User, error) {
	return u.user.GetUsers()
}

func (u *UserService) UpdateUser(user *domain.User) error {
	return u.user.UpdateUser(user)
}

func (u *UserService) DeleteUser(id int) error {
	return u.user.DeleteUser(id)
}
