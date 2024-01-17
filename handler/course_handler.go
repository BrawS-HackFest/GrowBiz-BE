package handler

import (
	"HackFest/models"
	"HackFest/service"
	"HackFest/service/courses"
	"HackFest/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type CourseHandler struct {
	courseService courses.CourseService
	userService   service.UserService
}

func NewCourseHandler(courseService courses.CourseService, userService service.UserService) *CourseHandler {
	return &CourseHandler{
		courseService: courseService,
		userService:   userService,
	}
}

func (ch *CourseHandler) FindAll(c *gin.Context) {
	courses, err := ch.courseService.FindAll()
	if err != nil {
		utils.HttpInternalError(c, "Can't get courses", err)
		return
	}

	var fixCourses []models.CourseResponse
	for _, course := range courses {
		data := models.CourseResponse{
			ID:     course.ID,
			Pict:   course.Link,
			Name:   course.Name,
			Buyer:  course.Buyer,
			Price:  course.Price,
			Rating: course.Rating,
		}
		fixCourses = append(fixCourses, data)
	}

	utils.HttpSuccess(c, "Success get courses", fixCourses)
}

func (ch *CourseHandler) FindByID(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	result, err := ch.courseService.FindByID(uint(id))
	if err != nil {
		utils.HttpInternalError(c, "Can't get courses", err)
		return
	}
	var reviews []models.ReviewResult
	for _, review := range result.Reviews {
		user, _ := ch.userService.FindByID(review.UserId)
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
		reviews = append(reviews, data)
	}

	bab := utils.SplitBabString(result.Bab)

	fixResult := models.CourseResponseByID{
		ID:     result.ID,
		Name:   result.Name,
		Desc:   result.Desc,
		Price:  result.Price,
		Buyer:  result.Buyer,
		Rating: result.Rating,
		Bab:    bab,
		Pict:   result.Link,
		Review: reviews,
	}
	utils.HttpSuccess(c, "Success get courses", fixResult)
}

func (ch *CourseHandler) Create(c *gin.Context) {
	c1 := models.Course{
		Model:  gorm.Model{},
		Name:   "Dasar Digital Marketing",
		Desc:   "Melalui program Dasar Digital Msrketing, kami ingin membuka kesempatan untuk siapa saja yang memulai karir di industri digital. Oleh karena itu, untuk mengikuti program ini, peserta program tidak harus memiliki background marketing ataupun digital marketing sebelumnya.",
		Price:  50000,
		Buyer:  0,
		Rating: 0,
		Bab:    "Trailer;Market Research;Wordpress;Email Marketing;Youtube Marketing;Linkedin Marketing",
		Link:   "https://kpkmxnicmhvpqmxywspm.supabase.co/storage/v1/object/public/picture/digitalmarketing.jpg",
	}
	c2 := models.Course{
		Model:  gorm.Model{},
		Name:   "Membangun Usaha kuliner",
		Desc:   "Course \"Membangun Usaha Kuliner\" dirancang untuk memberikan pemahaman mendalam tentang langkah-langkah yang diperlukan dalam mendirikan dan mengelola bisnis kuliner. Peserta akan diajarkan konsep-konsep kunci seperti perencanaan bisnis, pengembangan menu, manajemen operasional, pemasaran, dan aspek hukum yang terkait dengan industri kuliner.",
		Price:  45000,
		Buyer:  0,
		Rating: 0,
		Bab:    "Bab 1;Bab 2;Bab 3;Bab 4;Bab 5",
		Link:   "https://kpkmxnicmhvpqmxywspm.supabase.co/storage/v1/object/public/picture/culiner.jpg?t=2024-01-09T14%3A31%3A08.047Z",
	}
	c3 := models.Course{
		Model:  gorm.Model{},
		Name:   "Langkah Awal Menjadi Influencer",
		Desc:   "Course \"Langkah Awal Menjadi Influencer\" dirancang untuk membimbing peserta dalam memahami dan melangkah dalam dunia menjadi seorang influencer. Materi kursus ini mencakup berbagai aspek, mulai dari membangun kehadiran online hingga strategi pemasaran personal.",
		Price:  60000,
		Buyer:  0,
		Rating: 0,
		Bab:    "Bab 1;Bab 2;Bab 3;Bab 4;Bab 5",
		Link:   "https://kpkmxnicmhvpqmxywspm.supabase.co/storage/v1/object/public/picture/influencer.jpg",
	}

	c4 := models.Course{
		Model:  gorm.Model{},
		Name:   "Memulai Usaha Nail Art di Usia Muda",
		Desc:   "Peserta akan belajar tentang berbagai teknik nail art, tren terkini, dan cara menghadirkan keunikan dalam desain kuku mereka. Selain itu, kursus ini akan membahas aspek-aspek bisnis seperti perencanaan keuangan, manajemen inventaris, dan penentuan harga layanan.",
		Price:  70000,
		Buyer:  0,
		Rating: 0,
		Bab:    "Bab 1;Bab 2;Bab 3;Bab 4;Bab 5",
		Link:   "https://kpkmxnicmhvpqmxywspm.supabase.co/storage/v1/object/public/picture/Untitled.png?t=2024-01-09T14%3A43%3A46.788Z",
	}

	if _, err := ch.courseService.Create(c1); err != nil {
		utils.HttpInternalError(c, "Can't create courses", err)
		return
	}
	if _, err := ch.courseService.Create(c2); err != nil {
		utils.HttpInternalError(c, "Can't create courses", err)
		return
	}
	if _, err := ch.courseService.Create(c3); err != nil {
		utils.HttpInternalError(c, "Can't create courses", err)
		return
	}
	if _, err := ch.courseService.Create(c4); err != nil {
		utils.HttpInternalError(c, "Can't create courses", err)
		return
	}
}
