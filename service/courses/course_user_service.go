package courses

import (
	"HackFest/models"
	"HackFest/repository/courses"
)

type CourseUserService interface {
	Create(data models.CourseUserPost) (models.CourseUser, error)
	Update(courseID uint, userID string) error
	FindByCourseIDAndUserID(courseID uint, userID string) (models.CourseUser, error)
	FindByUserID(id string) ([]models.CourseUser, error)
}

type courseUserService struct {
	courseUserRepository courses.CourseUserRepository
}

func NewCourseUserService(courseUserRepository courses.CourseUserRepository) CourseUserService {
	return &courseUserService{courseUserRepository}
}

func (r *courseUserService) Create(data models.CourseUserPost) (models.CourseUser, error) {
	result := models.CourseUser{
		CourseID: data.CourseID,
		UserID:   data.UserID,
		IsRated:  false,
	}
	result, err := r.courseUserRepository.Create(result)
	if err != nil {
		return models.CourseUser{}, err
	}
	return result, nil
}

func (r *courseUserService) Update(courseID uint, userID string) error {
	data, err := r.courseUserRepository.FindByCourseIDAndUserID(courseID, userID)
	if err != nil {
		return err
	}
	data.IsRated = true
	if err := r.courseUserRepository.Update(courseID, userID, data); err != nil {
		return err
	}
	return nil
}

func (r *courseUserService) FindByCourseIDAndUserID(courseID uint, userID string) (models.CourseUser, error) {
	result, err := r.courseUserRepository.FindByCourseIDAndUserID(courseID, userID)
	if err != nil {
		return models.CourseUser{}, err
	}
	return result, nil
}

func (r *courseUserService) FindByUserID(id string) ([]models.CourseUser, error) {
	data, err := r.courseUserRepository.FindByUserID(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}
