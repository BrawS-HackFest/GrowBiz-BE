package repository

import (
	"HackFest/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user models.User) (models.User, error)
	FindByID(id string) (models.User, error)
	FindAll() ([]models.User, error)
	UpdateUser(id string, categories []models.Category, user models.User) error
	GetProfile(id string) (models.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (u *userRepository) Create(user models.User) (models.User, error) {
	if err := u.db.Create(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (u *userRepository) FindByID(id string) (models.User, error) {
	var user models.User
	if err := u.db.First(&user, "id = ?", id).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func (u *userRepository) FindAll() ([]models.User, error) {
	var data []models.User
	if err := u.db.Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}

func (u *userRepository) UpdateUser(id string, categories []models.Category, user models.User) error {
	if err := u.db.Model(&user).Where("id = ?", id).Updates(user).Error; err != nil {
		return err
	}
	if err := u.db.Model(&user).Where("id = ?", id).Association("Categories").Append(&categories); err != nil {
		return err
	}
	return nil
}

func (u *userRepository) GetProfile(id string) (models.User, error) {
	var user models.User
	if err := u.db.First(&user, "id = ?", id).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}
