package handler

import (
	"HackFest/models"
	"HackFest/service/courses"
	"HackFest/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type MaterialHandler struct {
	materialService courses.MaterialService
}

func NewMaterialHandler(materialService courses.MaterialService) *MaterialHandler {
	return &MaterialHandler{materialService}
}

func (mh *MaterialHandler) Create(c *gin.Context) {
	m1 := models.CourseMaterial{
		Model:    gorm.Model{},
		Title:    "Trailer",
		Content:  "https://www.youtube.com/watch?v=icg9b6itW6I&pp=ygUWcGFzYXIgZGlnaWF0IG1hcmtldGluZw%3D%3D",
		CourseID: 1,
	}

	m2 := models.CourseMaterial{
		Model:    gorm.Model{},
		Title:    "Market Research",
		Content:  "https://www.youtube.com/watch?v=b-hDg7699S0&pp=ygUPbWFya2V0IHJlc2VhcmNo",
		CourseID: 1,
	}

	m3 := models.CourseMaterial{
		Model:    gorm.Model{},
		Title:    "Wordpress",
		Content:  "https://www.youtube.com/watch?v=71EZb94AS1k&pp=ygUJd29yZHByZXNz",
		CourseID: 1,
	}

	m4 := models.CourseMaterial{
		Model:    gorm.Model{},
		Title:    "Bab 1",
		Content:  "Ini Meruapakan materi mengenai bab 1",
		CourseID: 2,
	}

	m5 := models.CourseMaterial{
		Model:    gorm.Model{},
		Title:    "Bab 2",
		Content:  "Ini bab 2",
		CourseID: 2,
	}

	m6 := models.CourseMaterial{
		Model:    gorm.Model{},
		Title:    "Bab 3",
		Content:  "Ini bab 3",
		CourseID: 2,
	}

	if _, err := mh.materialService.Create(m1); err != nil {
		utils.HttpInternalError(c, "Can't create material", err)
		return
	}
	if _, err := mh.materialService.Create(m2); err != nil {
		utils.HttpInternalError(c, "Can't create material", err)
		return
	}
	if _, err := mh.materialService.Create(m3); err != nil {
		utils.HttpInternalError(c, "Can't create material", err)
		return
	}
	if _, err := mh.materialService.Create(m4); err != nil {
		utils.HttpInternalError(c, "Can't create material", err)
		return
	}
	if _, err := mh.materialService.Create(m5); err != nil {
		utils.HttpInternalError(c, "Can't create material", err)
		return
	}
	if _, err := mh.materialService.Create(m6); err != nil {
		utils.HttpInternalError(c, "Can't create material", err)
		return
	}
}

type CourseByID struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
}

func (mh *MaterialHandler) FindByCourseID(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	result, err := mh.materialService.FindByCourseID(uint(id))
	if err != nil {
		utils.HttpInternalError(c, "Can't get materials", err)
		return
	}

	var data []CourseByID
	for _, v := range result {
		data = append(data, CourseByID{
			ID:    v.ID,
			Title: v.Title,
		})
	}
	utils.HttpSuccess(c, "Success get materials", data)
}

func (mh *MaterialHandler) FindByID(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	data, err := mh.materialService.FindByID(uint(id))
	if err != nil {
		utils.HttpInternalError(c, "Can't get material", err)
		return
	}
	utils.HttpSuccess(c, "Success get material", data)
}
