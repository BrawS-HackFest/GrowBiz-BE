package courses

import (
	"HackFest/models"
	"HackFest/repository/courses"
)

type CourseService interface {
	Create(data models.Course) (models.Course, error)
	FindByID(id uint) (models.Course, error)
	FindAll() ([]models.Course, error)
}

type courseService struct {
	courseRepo courses.CourseRepository
}

func NewCourseService(courseRepo courses.CourseRepository) CourseService {
	return &courseService{courseRepo}
}

func (c *courseService) Create(data models.Course) (models.Course, error) {
	create, err := c.courseRepo.Create(data)
	if err != nil {
		return models.Course{}, err
	}
	return create, nil
}

func (c *courseService) FindByID(id uint) (models.Course, error) {
	data, err := c.courseRepo.FindByID(id)
	if err != nil {
		return models.Course{}, err
	}
	return data, nil
}

func (c *courseService) FindAll() ([]models.Course, error) {
	return c.courseRepo.FindAll()
}
