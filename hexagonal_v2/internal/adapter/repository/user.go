package repository

import (
	"hexagonal_v2/internal/core/domain"
	"hexagonal_v2/internal/core/port"

	"gorm.io/gorm"
)

type userRepository struct {
	Db *gorm.DB
}

func NewUserRepository(Db *gorm.DB) port.UserRepository {
	Db.AutoMigrate(&domain.User{})
	return &userRepository{Db: Db}
}

func (u *userRepository) CreateUser(user *domain.User) error {
	return u.Db.Create(user).Error
}

func (u *userRepository) GetUserByID(id int64) (*domain.User, error) {
	var user domain.User
	err := u.Db.Preload("Airpods").First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *userRepository) GetUsers() ([]*domain.User, error) {
	var users []*domain.User
	err := u.Db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *userRepository) UpdateUser(user *domain.User) error {
	return u.Db.Save(user).Error
}

func (u *userRepository) DeleteUser(id int64) error {
	return u.Db.Delete(&domain.User{}, id).Error
}
