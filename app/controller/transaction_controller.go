package controller

import (
	"kasir-api-gin/domains/entity"
	"kasir-api-gin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionController struct {
	service service.TransactionService
}

func NewTransactionController(service service.TransactionService) transactionController {
	return transactionController{
		service: service,
	}
}

func (controller transactionController) PostTransaction(ctx *gin.Context) {
	var inputTransaction entity.Transaction
	if err := ctx.ShouldBind(&inputTransaction); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	user_id, _ := ctx.Get("user_id")
	inputTransaction.UserID = user_id.(uint)
	transaction, err := controller.service.Save(inputTransaction)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message":     "Transaction successfully created",
		"transaction": transaction,
	})
}

func (controller transactionController) GetTransaction(ctx *gin.Context) {
	stringId := ctx.Param("transaction_id")
	transaction, err := controller.service.GetById(stringId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"transaction": transaction,
	})
}
