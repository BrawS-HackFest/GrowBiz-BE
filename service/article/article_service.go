package article

import (
	"HackFest/models"
	"HackFest/repository"
	"HackFest/repository/article"
)

type ArticleService interface {
	Create(article models.Article) error
	FindAll() ([]models.Article, error)
	FindByID(id uint) (models.ArticleByIDResponse, error)
}

type articleService struct {
	articleRepository article.ArticleRepository
	userRepository    repository.UserRepository
}

func NewArticleService(articleRepository article.ArticleRepository, userRepository repository.UserRepository) ArticleService {
	return &articleService{
		articleRepository,
		userRepository,
	}
}

func (a *articleService) Create(article models.Article) error {
	return a.articleRepository.Create(article)
}

func (a *articleService) FindAll() ([]models.Article, error) {
	return a.articleRepository.FindAll()
}

func (a *articleService) FindByID(id uint) (models.ArticleByIDResponse, error) {
	data, err := a.articleRepository.FindByID(id)
	if err != nil {
		return models.ArticleByIDResponse{}, err
	}
	var comments []models.Comment

	for _, comment := range data.Comments {
		user, _ := a.userRepository.FindByID(comment.UserID)
		comments = append(comments, models.Comment{
			ID:        comment.ID,
			UpdatedAt: comment.UpdatedAt,
			Comment:   comment.Comment,
			ArticleID: comment.ArticleID,
			User: models.UserReview{
				ID:       user.Id,
				Username: user.Username,
			},
		})
	}
	result := models.ArticleByIDResponse{
		ID:          data.ID,
		UpdatedAt:   data.UpdatedAt,
		Title:       data.Title,
		Pict:        data.Pict,
		Description: data.Description,
		Comments:    comments,
	}

	return result, nil
}
