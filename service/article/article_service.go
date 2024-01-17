package article

import (
	"HackFest/models"
	"HackFest/repository/article"
)

type ArticleService interface {
	Create(article models.Article) error
	FindAll() ([]models.Article, error)
	FindByID(id uint) (models.Article, error)
}

type articleService struct {
	articleRepository article.ArticleRepository
}

func NewArticleService(articleRepository article.ArticleRepository) ArticleService {
	return &articleService{
		articleRepository,
	}
}

func (a *articleService) Create(article models.Article) error {
	return a.articleRepository.Create(article)
}

func (a *articleService) FindAll() ([]models.Article, error) {
	return a.articleRepository.FindAll()
}

func (a *articleService) FindByID(id uint) (models.Article, error) {
	return a.articleRepository.FindByID(id)
}
