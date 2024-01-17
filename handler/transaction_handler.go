package handler

import (
	"HackFest/models"
	"HackFest/service"
	"HackFest/utils"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"strconv"
)

type TransactionHandler struct {
	transactionService service.TransactionService
}

func NewTransactionHandler(transactionService service.TransactionService) *TransactionHandler {
	return &TransactionHandler{transactionService}
}

func (th *TransactionHandler) Create(c *gin.Context) {
	id := c.MustGet("userID").(string)
	var transaction models.TransactionPost
	if err := c.ShouldBindJSON(&transaction); err != nil {
		utils.HttpFailOrError(c, 400, "failed to bind json", err)
		return
	}
	create, err := th.transactionService.Create(id, transaction)
	if err != nil {
		utils.HttpInternalError(c, "Failed to create transaction", err)
		return
	}

	if create.UserID == "" {
		utils.HttpFailOrError(c, 400, "Already bought the courses", err)
		return
	}
	utils.HttpSuccess(c, "Transaction created successfully", create)
}

func (th *TransactionHandler) FindAll(c *gin.Context) {
	data, err := th.transactionService.FindAll()
	if err != nil {
		utils.HttpInternalError(c, "Failed to fetch transactions", err)
		return
	}
	utils.HttpSuccess(c, "Transactions fetched successfully", data)
}

func (th *TransactionHandler) FindByID(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.Atoi(idStr)
	data, err := th.transactionService.FindByID(uint(id))
	if err != nil {
		utils.HttpInternalError(c, "Failed to fetch transaction", err)
		return
	}

	result := models.TransactionByID{
		ID:       data.ID,
		Method:   data.Method,
		Amount:   data.Amount,
		VANumber: data.VaNumber,
		OrderID:  data.OrderID,
		Status:   data.Status,
	}

	utils.HttpSuccess(c, "Transaction fetched successfully", result)
}

func (th *TransactionHandler) FindByUserID(c *gin.Context) {
	idUser := c.MustGet("userID").(string)
	data, err := th.transactionService.FindByUserID(idUser)
	if err != nil {
		utils.HttpInternalError(c, "Failed to fetch transactions", err)
		return
	}
	utils.HttpSuccess(c, "Transactions fetched successfully", data)
}

func (th *TransactionHandler) Update(c *gin.Context) {
	var notifPayload map[string]interface{}

	err := json.NewDecoder(c.Request.Body).Decode(&notifPayload)
	if err != nil {
		utils.HttpFailOrError(c, 404, "failed to decode payload", err)
		return
	}
	orderID, exist := notifPayload["order_id"].(string)
	if !exist {
		utils.HttpFailOrError(c, 404, "order id not found", err)
		return
	}
	var data models.Transaction
	data, err = th.transactionService.Update(orderID)
	if err != nil {
		utils.HttpInternalError(c, "Failed to update transaction", err)
		return
	}

	utils.HttpSuccess(c, "Transaction updated successfully", data)
}
