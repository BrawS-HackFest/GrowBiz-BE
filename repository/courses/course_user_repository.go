package courses

import (
	"HackFest/models"
	"gorm.io/gorm"
)

type CourseUserRepository interface {
	Create(data models.CourseUser) (models.CourseUser, error)
	Update(courseID uint, userID string, data models.CourseUser) error
	FindByCourseIDAndUserID(courseID uint, userID string) (models.CourseUser, error)
	FindByUserID(id string) ([]models.CourseUser, error)
}

type courseUserRepository struct {
	db *gorm.DB
}

func NewCourseUserRepository(db *gorm.DB) CourseUserRepository {
	return &courseUserRepository{db}
}

func (r *courseUserRepository) Create(data models.CourseUser) (models.CourseUser, error) {
	if err := r.db.Create(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func (r *courseUserRepository) Update(courseID uint, userID string, data models.CourseUser) error {
	if err := r.db.Where("course_id = ? AND user_id = ?", courseID, userID).Save(&data).Error; err != nil {
		return err
	}
	return nil
}

func (r *courseUserRepository) FindByCourseIDAndUserID(courseID uint, userID string) (models.CourseUser, error) {
	var data models.CourseUser
	if err := r.db.Where("course_id = ? AND user_id = ?", courseID, userID).First(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func (r *courseUserRepository) FindByUserID(id string) ([]models.CourseUser, error) {
	var data []models.CourseUser
	if err := r.db.Where("user_id = ?", id).Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}
