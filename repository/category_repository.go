package repository

import (
	"HackFest/models"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Find(id []uint) ([]models.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{db}
}

func (c *categoryRepository) Find(id []uint) ([]models.Category, error) {
	var categories []models.Category
	if err := c.db.Find(&categories, id).Error; err != nil {
		return nil, err
	}
	return categories, nil
}
