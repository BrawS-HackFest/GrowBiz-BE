package courses

import (
	"HackFest/models"
	"gorm.io/gorm"
)

type ReviewRepository interface {
	Create(review models.Review) (models.Review, error)
	Update(id uint, review models.Review) (models.Review, error)
	FindByCourseID(idCourse uint) ([]models.Review, error)
}

type reviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) ReviewRepository {
	return &reviewRepository{db}
}

func (r *reviewRepository) Create(review models.Review) (models.Review, error) {
	err := r.db.Create(&review).Error
	if err != nil {
		return models.Review{}, err
	}
	return review, nil
}

func (r *reviewRepository) Update(id uint, review models.Review) (models.Review, error) {
	err := r.db.Model(&review).Where("id = ?", id).Updates(review).Error
	if err != nil {
		return models.Review{}, err
	}
	return review, nil
}

func (r *reviewRepository) FindByCourseID(idCourse uint) ([]models.Review, error) {
	var review []models.Review
	if err := r.db.Where("course_id = ?", idCourse).Find(&review).Error; err != nil {
		return nil, err
	}
	return review, nil
}
