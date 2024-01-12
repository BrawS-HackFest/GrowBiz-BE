package service

import (
	"HackFest/config/mdtrans"
	"HackFest/models"
	"HackFest/repository"
	"HackFest/utils"
	"gorm.io/gorm"
	"log"
)

type TransactionService interface {
	Create(userID string, transaction models.TransactionPost) (models.Transaction, error)
	FindAll() ([]models.Transaction, error)
	FindByID(id uint) (models.Transaction, error)
	FindByUserID(id string) ([]models.Transaction, error)
	Update(orderID string) (models.Transaction, error)
}

type transactionService struct {
	transactionRepository repository.TransactionRepository
	userRepository        repository.UserRepository
	courseRepository      repository.CourseRepository
	courseUserRepository  repository.CourseUserRepository
	midtransClient        *mdtrans.MdtClient
}

func NewTransactionService(
	transactionRepository repository.TransactionRepository,
	userRepository repository.UserRepository,
	courseRepository repository.CourseRepository,
	courseUserRepository repository.CourseUserRepository,
	midtransClient *mdtrans.MdtClient) TransactionService {
	return &transactionService{
		transactionRepository: transactionRepository,
		userRepository:        userRepository,
		courseRepository:      courseRepository,
		courseUserRepository:  courseUserRepository,
		midtransClient:        midtransClient,
	}
}

func (ts *transactionService) Create(userID string, transaction models.TransactionPost) (models.Transaction, error) {
	user, err := ts.userRepository.FindByID(userID)
	if err != nil {
		return models.Transaction{}, err
	}
	course, err := ts.courseRepository.FindByID(uint(transaction.CourseID))
	if err != nil {
		return models.Transaction{}, err
	}
	res, _ := ts.courseUserRepository.FindByCourseIDAndUserID(uint(transaction.CourseID), userID)
	if res.UserID != "" {
		return models.Transaction{}, err
	}

	orderID := utils.GenerateId()

	midtrans, err := ts.midtransClient.CreateOrder(orderID, transaction.Amount, uint(transaction.CourseID),
		course.Name, user.Email, user.Username, transaction.Method)

	data := models.Transaction{
		Model:         gorm.Model{},
		OrderID:       orderID,
		TransactionID: midtrans.TransactionID,
		Amount:        transaction.Amount,
		Method:        transaction.Method,
		VaNumber:      midtrans.VaNumbers[0].VANumber,
		Status:        "pending",
		CourseID:      transaction.CourseID,
		UserID:        userID,
	}

	create, err := ts.transactionRepository.Create(data)
	if err != nil {
		return models.Transaction{}, err
	}
	return create, nil
}

func (ts *transactionService) FindAll() ([]models.Transaction, error) {
	transactions, err := ts.transactionRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (ts *transactionService) FindByID(id uint) (models.Transaction, error) {
	data, err := ts.transactionRepository.FindByID(id)
	if err != nil {
		return models.Transaction{}, err
	}
	return data, nil
}

func (ts *transactionService) FindByUserID(id string) ([]models.Transaction, error) {
	transactions, err := ts.transactionRepository.FindByUserID(id)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (ts *transactionService) Update(orderID string) (models.Transaction, error) {
	data, err := ts.transactionRepository.FindByOrderID(orderID)
	if err != nil {
		return models.Transaction{}, err
	}

	transStatus, err := ts.midtransClient.NotifHandler(orderID)
	if err != nil {
		return models.Transaction{}, err
	} else {
		if transStatus != nil {
			if transStatus.TransactionStatus == "capture" {
				if transStatus.FraudStatus == "challenge" {
					data.Status = "challenge"
				} else if transStatus.FraudStatus == "accept" {
					data.Status = "success"
				}
			} else if transStatus.TransactionStatus == "settlement" {
				data.Status = "success"
			} else if transStatus.TransactionStatus == "deny" {
				data.Status = "deny"
			} else if transStatus.TransactionStatus == "cancel" || transStatus.TransactionStatus == "expire" {
				data.Status = "failure"
			} else if transStatus.TransactionStatus == "pending" {
				data.Status = "pending"
			}
		}
	}

	if data.Status == "success" {
		courseUser := models.CourseUser{
			CourseID: uint(data.CourseID),
			UserID:   data.UserID,
			IsRated:  false,
		}
		_, err = ts.courseUserRepository.Create(courseUser)
		if err != nil {
			log.Println("=========================\n", err, "\n2=========================")
			return models.Transaction{}, err
		}
	}

	var result models.Transaction
	result, err = ts.transactionRepository.Update(data)
	if err != nil {
		log.Println("=========================\n", err, "\n3=========================")
		return models.Transaction{}, err
	}
	return result, nil
}
