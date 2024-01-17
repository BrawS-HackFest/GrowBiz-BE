package service

import (
	"HackFest/models"
	"HackFest/repository"
)

type ArticleService interface {
	Create(article models.Article) error
	FindAll() ([]models.Article, error)
	FindByID(id uint) (models.Article, error)
}

type articleService struct {
	articleRepository repository.ArticleRepository
}

func NewArticleService(articleRepository repository.ArticleRepository) ArticleService {
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
