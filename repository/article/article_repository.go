package article

import (
	"HackFest/models"
	"gorm.io/gorm"
	"log"
)

type ArticleRepository interface {
	Create(article models.Article) error
	FindAll() ([]models.Article, error)
	FindByID(id uint) (models.Article, error)
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{db}
}

func (a *articleRepository) Create(article models.Article) error {
	if err := a.db.Create(&article).Error; err != nil {
		return err
	}
	return nil
}

func (a *articleRepository) FindAll() ([]models.Article, error) {
	var data []models.Article
	err := a.db.Find(&data).Error
	if err != nil {
		log.Println("Error FindAll article: ", err)
		return nil, err
	}
	return data, nil
}

func (a *articleRepository) FindByID(id uint) (models.Article, error) {
	var data models.Article
	err := a.db.First(&data, id).Error
	if err != nil {
		log.Println("Error FindByID article: ", err)
		return data, err
	}
	return data, nil
}
