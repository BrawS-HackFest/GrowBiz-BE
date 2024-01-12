package handler

import (
	"HackFest/models"
	"HackFest/service"
	"HackFest/utils"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ReviewHandler struct {
	reviewService service.ReviewService
	userService   service.UserService
}

func NewReviewHandler(reviewService service.ReviewService, userService service.UserService) *ReviewHandler {
	return &ReviewHandler{
		reviewService: reviewService,
		userService:   userService,
	}
}

func (h *ReviewHandler) Create(c *gin.Context) {

}

func (h *ReviewHandler) FindByCourseID(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	result, err := h.reviewService.FindByCourseID(uint(id))
	if err != nil {
		utils.HttpInternalError(c, "Can't get reviews", err)
		return
	}
	var datas []models.ReviewResult
	for _, review := range result {
		user, _ := h.userService.FindByID(review.UserId)
		data := models.ReviewResult{
			ID:        review.ID,
			UpdatedAt: review.UpdatedAt,
			Comment:   review.Comment,
			Rating:    review.Rating,
			CourseID:  review.CourseId,
			User: models.UserReview{
				ID:       user.Id,
				Username: user.Username,
			},
		}
		datas = append(datas, data)
	}

	utils.HttpSuccess(c, "Success get reviews", datas)
}

func (h *ReviewHandler) Update(c *gin.Context) {
	
}
