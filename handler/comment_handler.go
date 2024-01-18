package handler

import (
	"HackFest/models"
	"HackFest/service/article"
	"HackFest/utils"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

type CommentHandler struct {
	commentService article.CommentService
}

func NewCommentHandler(commentService article.CommentService) *CommentHandler {
	return &CommentHandler{commentService}
}

func (ch *CommentHandler) Create(c *gin.Context) {
	articleIDStr := c.Param("articleID")
	articleID, _ := strconv.Atoi(articleIDStr)
	userID := c.MustGet("userID").(string)

	var comment models.CommentPost
	if err := c.ShouldBindJSON(&comment); err != nil {
		utils.HttpFailOrError(c, 400, "Bad request", err)
		return
	}
	post, err := ch.commentService.Create(comment.Comment, uint(articleID), userID)
	if err != nil {
		utils.HttpInternalError(c, "Can't create comment", err)
		return
	}
	log.Println("=====================\n", post.ID, "\n", post.Comment, "\n=====================")
	utils.HttpSuccess(c, "Success create comment", post)
}

func (ch *CommentHandler) FindByArticleID(c *gin.Context) {
	articleIDStr := c.Param("articleID")
	articleID, _ := strconv.Atoi(articleIDStr)
	data, err := ch.commentService.FindByArticleID(uint(articleID))
	if err != nil {
		utils.HttpInternalError(c, "Can't get comments", err)
		return
	}
	utils.HttpSuccess(c, "Success get comments", data)
}

func (ch *CommentHandler) FindByID(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	data, err := ch.commentService.FindByID(uint(id))
	if err != nil {
		utils.HttpInternalError(c, "Can't get comment", err)
		return
	}
	utils.HttpSuccess(c, "Success get comment", data)
}

func (ch *CommentHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)

	var comment models.CommentPost
	if err := c.ShouldBindJSON(&comment); err != nil {
		utils.HttpFailOrError(c, 400, "Bad request", err)
		return
	}
	if err := ch.commentService.Update(uint(id), comment.Comment); err != nil {
		utils.HttpInternalError(c, "Can't update comment", err)
		return
	}
	utils.HttpSuccess(c, "Success update comment", nil)
}

func (ch *CommentHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	if err := ch.commentService.Delete(uint(id)); err != nil {
		utils.HttpInternalError(c, "Can't delete comment", err)
		return
	}
	utils.HttpSuccess(c, "Success delete comment", nil)
}
