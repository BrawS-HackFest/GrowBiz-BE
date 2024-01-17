package courses

import (
	"HackFest/models"
	"gorm.io/gorm"
)

type CourseRepository interface {
	Create(data models.Course) (models.Course, error)
	FindByID(id uint) (models.Course, error)
	FindAll() ([]models.Course, error)
}

type courseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{db}
}

func (c *courseRepository) Create(data models.Course) (models.Course, error) {
	if err := c.db.Create(&data).Error; err != nil {
		return models.Course{}, err
	}
	return data, nil
}

func (c *courseRepository) FindByID(id uint) (models.Course, error) {
	var course models.Course
	if err := c.db.Model(&course).Preload("Reviews").First(&course, id).Error; err != nil {
		return models.Course{}, err
	}
	return course, nil
}

func (c *courseRepository) FindAll() ([]models.Course, error) {
	var data []models.Course
	if err := c.db.Model(&models.Course{}).Order("buyer desc").Find(&data).Error; err != nil {
		return nil, err
	}
	return data, nil
}
