package handler

import (
	"HackFest/models"
	"HackFest/service"
	"HackFest/utils"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{userService}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var data models.UserCreate
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.HttpFailOrError(c, 400, "Failed to bind JSON", err)
		return
	}
	user := models.UserCreate{
		Id:    data.Id,
		Email: data.Email,
	}
	create, err := h.userService.Create(user)
	if err != nil {
		utils.HttpInternalError(c, "Failed to create user", err)
		return
	}
	utils.HttpSuccess(c, "Successfully created user", create)
}

func (h *UserHandler) FindUser(c *gin.Context) {
	id := c.MustGet("userID").(string)
	user, err := h.userService.FindByID(id)
	if err != nil {
		utils.HttpInternalError(c, "Failed to find user", err)
		return
	}
	utils.HttpSuccess(c, "Successfully found user", user)
}

func (h *UserHandler) FindAllUser(c *gin.Context) {
	users, err := h.userService.FindAll()
	if err != nil {
		utils.HttpInternalError(c, "Failed to find users", err)
		return
	}
	utils.HttpSuccess(c, "Successfully found users", users)
}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	//id := c.MustGet("userID").(string)
	var data models.UserUpdate
	if err := c.ShouldBindJSON(&data); err != nil {
		utils.HttpFailOrError(c, 400, "Failed to bind JSON", err)
		return
	}
	err := h.userService.UpdateUser("dZtS5UZUC3MX6GABFlxhEluj4XH2", data.Number, data.Username, data.Category)
	if err != nil {
		utils.HttpInternalError(c, "Failed to update user", err)
		return
	}
	utils.HttpSuccess(c, "Successfully updated user", nil)
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	id := c.MustGet("userID").(string)
	user, err := h.userService.GetProfile(id)
	if err != nil {
		utils.HttpInternalError(c, "Failed to get profile", err)
		return
	}
	utils.HttpSuccess(c, "Successfully get profile", user)
}
