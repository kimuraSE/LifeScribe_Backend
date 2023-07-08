package repository

import (
	"LifeScribe_Backend/internal/api/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
	DeleteUser(userId uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUserByEmail(user *model.User, email string) error {
	if err := r.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) CreateUser(user *model.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) DeleteUser(userId uint) error {
	if err := r.db.Where("id=?", userId).Delete(&model.User{}).Error; err != nil {
		return err
	}
	return nil
}
