package repository

import (
	"HackFest/models"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Create(transaction models.Transaction) (models.Transaction, error)
	FindAll() ([]models.Transaction, error)
	FindByID(id uint) (models.Transaction, error)
	FindByUserID(id string) ([]models.Transaction, error)
	FindByOrderID(id string) (models.Transaction, error)
	Update(transaction models.Transaction) (models.Transaction, error)
}

type transactionRepository struct {
	db *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepository{db}
}

func (tr *transactionRepository) Create(transaction models.Transaction) (models.Transaction, error) {
	if err := tr.db.Create(&transaction).Error; err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (tr *transactionRepository) FindAll() ([]models.Transaction, error) {
	var data []models.Transaction
	if err := tr.db.Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func (tr *transactionRepository) FindByID(id uint) (models.Transaction, error) {
	var data models.Transaction
	if err := tr.db.First(&data, id).Error; err != nil {
		return data, err
	}
	return data, nil
}

func (tr *transactionRepository) FindByUserID(id string) ([]models.Transaction, error) {
	var data []models.Transaction
	if err := tr.db.Where("user_id = ?", id).Where("status = ?", "pending").Find(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}

func (tr *transactionRepository) FindByOrderID(id string) (models.Transaction, error) {
	var data models.Transaction
	if err := tr.db.Where("order_id = ?", id).First(&data).Error; err != nil {
		return data, err
	}
	return data, nil
}
func (tr *transactionRepository) Update(transaction models.Transaction) (models.Transaction, error) {
	if err := tr.db.Where("id = ?", transaction.ID).Save(&transaction).Error; err != nil {
		return transaction, err
	}
	return transaction, nil
}
