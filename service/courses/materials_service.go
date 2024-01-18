package courses

import (
	"HackFest/models"
	"HackFest/repository/courses"
)

type MaterialService interface {
	Create(material models.CourseMaterial) (models.CourseMaterial, error)
	FindByCourseID(id uint) ([]models.CourseMaterial, error)
	FindByID(id uint) (models.CourseMaterial, error)
}

type materialService struct {
	materialRepository courses.MaterialRepository
}

func NewMaterialService(materialRepository courses.MaterialRepository) MaterialService {
	return &materialService{
		materialRepository,
	}
}

func (m *materialService) Create(material models.CourseMaterial) (models.CourseMaterial, error) {
	return m.materialRepository.Create(material)
}

func (m *materialService) FindByCourseID(id uint) ([]models.CourseMaterial, error) {
	return m.materialRepository.FindByCourseID(id)
}

func (m *materialService) FindByID(id uint) (models.CourseMaterial, error) {
	return m.materialRepository.FindByID(id)
}
