package service

import (
	"HackFest/models"
	"HackFest/repository"
	"errors"
	"log"
)

type UserService interface {
	Create(user models.UserCreate) (models.User, error)
	FindByID(id string) (models.User, error)
	FindAll() ([]models.User, error)
	UpdateUser(id, number, username string, categories []uint) error
	GetProfile(id string) (models.User, error)
}

type userService struct {
	userRepository     repository.UserRepository
	categoryRepository repository.CategoryRepository
}

func NewUserService(userRepository repository.UserRepository, categoryRepository repository.CategoryRepository) UserService {
	return &userService{
		userRepository,
		categoryRepository,
	}
}

func (u *userService) Create(user models.UserCreate) (models.User, error) {
	if user.Id == "" {
		log.Println("id empty")
		return models.User{}, errors.New("empty id")
	}
	if user.Email == "" {
		log.Println("email empty")
		return models.User{}, errors.New("empty email")
	}
	data := models.User{
		Id:       user.Id,
		Username: "",
		Email:    user.Email,
		Number:   "",
	}
	create, err := u.userRepository.Create(data)
	if err != nil {
		return models.User{}, err
	}
	return create, nil
}

func (u *userService) FindByID(id string) (models.User, error) {
	user, err := u.userRepository.FindByID(id)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u *userService) FindAll() ([]models.User, error) {
	user, err := u.userRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userService) UpdateUser(id, number, username string, categories []uint) error {
	var user models.User
	data, err := u.userRepository.FindByID(id)
	if err != nil {
		return err
	}
	user.Number = number
	user.Username = username
	if number == "" {
		data.Number = number
	}
	if username == "" {
		user.Username = username
	}

	category, err := u.categoryRepository.Find(categories)
	if err != nil {
		return err
	}

	err = u.userRepository.UpdateUser(id, category, data)
	if err != nil {
		return err
	}
	return nil
}

func (u *userService) GetProfile(id string) (models.User, error) {
	user, err := u.userRepository.GetProfile(id)
	if err != nil {
		return user, err
	}
	return user, nil
}
