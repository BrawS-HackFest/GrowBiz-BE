package article

import (
	"HackFest/models"
	"HackFest/repository/article"
	"gorm.io/gorm"
)

type CommentService interface {
	Create(comment string, articleID uint, userID string) (models.ArticleComment, error)
	FindByArticleID(id uint) ([]models.ArticleComment, error)
	FindByID(id uint) (models.ArticleComment, error)
	Update(id uint, comment string) error
	Delete(id uint) error
}

type commentService struct {
	commentRepository article.CommentRepository
}

func NewCommentService(commentRepository article.CommentRepository) CommentService {
	return &commentService{commentRepository}
}

func (c *commentService) Create(comment string, articleID uint, userID string) (models.ArticleComment, error) {
	data := models.ArticleComment{
		Model:     gorm.Model{},
		Comment:   comment,
		ArticleID: articleID,
		UserID:    userID,
	}

	if err := c.commentRepository.Create(data); err != nil {
		return models.ArticleComment{}, err
	}
	return models.ArticleComment{}, nil
}

func (c *commentService) FindByArticleID(id uint) ([]models.ArticleComment, error) {
	data, err := c.commentRepository.FindByArticleID(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *commentService) FindByID(id uint) (models.ArticleComment, error) {
	data, err := c.commentRepository.FindByID(id)
	if err != nil {
		return data, err
	}
	return data, nil
}

func (c *commentService) Update(id uint, comment string) error {
	data, err := c.commentRepository.FindByID(id)
	if err != nil {
		return err
	}

	data.Comment = comment
	if err = c.commentRepository.Update(id, data); err != nil {
		return err
	}
	return nil
}

func (c *commentService) Delete(id uint) error {
	if err := c.commentRepository.Delete(id); err != nil {
		return err
	}
	return nil
}
