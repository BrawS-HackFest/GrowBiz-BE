package courses

import (
	"HackFest/models"
	"gorm.io/gorm"
)

type MaterialRepository interface {
	Create(material models.CourseMaterial) (models.CourseMaterial, error)
	FindByCourseID(id uint) ([]models.CourseMaterial, error)
	FindByID(id uint) (models.CourseMaterial, error)
}

type materialRepository struct {
	db *gorm.DB
}

func NewMaterialRepository(db *gorm.DB) MaterialRepository {
	return &materialRepository{db}
}

func (m *materialRepository) Create(material models.CourseMaterial) (models.CourseMaterial, error) {
	if err := m.db.Create(&material).Error; err != nil {
		return material, err
	}
	return material, nil
}

func (m *materialRepository) FindByCourseID(id uint) ([]models.CourseMaterial, error) {
	var data []models.CourseMaterial
	err := m.db.Where("course_id = ?", id).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (m *materialRepository) FindByID(id uint) (models.CourseMaterial, error) {
	var data models.CourseMaterial
	err := m.db.First(&data, id).Error
	if err != nil {
		return data, err
	}
	return data, nil
}
