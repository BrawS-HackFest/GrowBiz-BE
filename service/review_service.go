package service

import (
	"HackFest/models"
	"HackFest/repository"
	"gorm.io/gorm"
)

type ReviewService interface {
	Create(userID string, courseID uint, review models.ReviewPost) (models.Review, error)
	Update(id uint, comment string) (models.Review, error)
	FindByCourseID(idCourse uint) ([]models.Review, error)
}

type reviewService struct {
	reviewRepository repository.ReviewRepository
}

func NewReviewService(reviewRepository repository.ReviewRepository) ReviewService {
	return &reviewService{reviewRepository}
}

func (r *reviewService) Create(userID string, courseID uint, review models.ReviewPost) (models.Review, error) {
	data := models.Review{
		Model:    gorm.Model{},
		Comment:  review.Comment,
		Rating:   review.Rating,
		UserId:   userID,
		CourseId: courseID,
	}
	result, err := r.reviewRepository.Create(data)
	if err != nil {
		return models.Review{}, err
	}
	return result, nil
}

func (r *reviewService) Update(id uint, comment string) (models.Review, error) {
	var review models.Review
	review.Comment = comment
	res, err := r.reviewRepository.Update(id, review)
	if err != nil {
		return models.Review{}, err
	}
	return res, nil
}

func (r *reviewService) FindByCourseID(idCourse uint) ([]models.Review, error) {
	data, err := r.reviewRepository.FindByCourseID(idCourse)
	if err != nil {
		return nil, err
	}
	return data, nil
}
