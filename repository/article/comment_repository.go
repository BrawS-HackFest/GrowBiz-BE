package article

import (
	"HackFest/models"
	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(comment models.ArticleComment) error
	FindByArticleID(id uint) ([]models.ArticleComment, error)
	FindByID(id uint) (models.ArticleComment, error)
	Update(comment models.ArticleComment) error
	Delete(id uint) error
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{db}
}

func (c *commentRepository) Create(comment models.ArticleComment) error {
	if err := c.db.Create(&comment).Error; err != nil {
		return err
	}
	return nil
}

func (c *commentRepository) FindByArticleID(id uint) ([]models.ArticleComment, error) {
	var data []models.ArticleComment
	if err := c.db.Where("article_id = ?", id).Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func (c *commentRepository) FindByID(id uint) (models.ArticleComment, error) {
	var data models.ArticleComment
	if err := c.db.First(&data, id).Error; err != nil {
		return models.ArticleComment{}, err
	}
	return models.ArticleComment{}, nil
}

func (c *commentRepository) Update(comment models.ArticleComment) error {
	if err := c.db.Model(&comment).Update("comment", comment.Comment).Error; err != nil {
		return err
	}
	return nil
}

func (c *commentRepository) Delete(id uint) error {
	if err := c.db.Delete(&models.ArticleComment{}, id).Error; err != nil {
		return err
	}
	return nil
}
