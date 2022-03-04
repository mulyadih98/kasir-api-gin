package service

import (
	"kasir-api-gin/domains/entity"
	"kasir-api-gin/repository"
)

type transactionService struct {
	transactionRepository repository.TransactionRepository
}

type TransactionService interface {
	Save(entity.Transaction) (interface{}, error)
}

func NewTransactionService(transactionRepo repository.TransactionRepository) TransactionService {
	if err := transactionRepo.Migrate(); err != nil {
		panic(err.Error())
	}
	return transactionService{
		transactionRepository: transactionRepo,
	}
}

func (service transactionService) Save(transaction entity.Transaction) (interface{}, error) {
	return service.transactionRepository.Save(transaction)
}
