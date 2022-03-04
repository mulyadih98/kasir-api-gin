package service

import (
	"kasir-api-gin/domains/entity"
	"kasir-api-gin/repository"
	"strconv"
)

type transactionService struct {
	transactionRepository repository.TransactionRepository
	productRepository     repository.ProductRepository
}

type TransactionService interface {
	Save(entity.Transaction) (interface{}, error)
}

func NewTransactionService(transactionRepo repository.TransactionRepository, product repository.ProductRepository) TransactionService {
	if err := transactionRepo.Migrate(); err != nil {
		panic(err.Error())
	}
	return transactionService{
		transactionRepository: transactionRepo,
		productRepository:     product,
	}
}

func (service transactionService) Save(transaction entity.Transaction) (interface{}, error) {
	transaction, err := service.transactionRepository.Save(transaction)
	if err != nil {
		return transaction, err
	}

	// Mengurangi stock product
	for _, value := range transaction.TransactionDetail {
		service.productRepository.ReduceStock(strconv.FormatUint(uint64(value.ProductID), 10), value.Qty)
	}

	return transaction.SavedTransaction(), nil
}
