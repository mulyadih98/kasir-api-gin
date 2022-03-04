package repository

import (
	"kasir-api-gin/domains/entity"

	"gorm.io/gorm"
)

type transactionRepositoryGorm struct {
	DB *gorm.DB
}

func NewTransactionRepositoryGorm(db *gorm.DB) TransactionRepository {
	return transactionRepositoryGorm{
		DB: db,
	}
}

func (repo transactionRepositoryGorm) Migrate() error {
	return repo.DB.AutoMigrate(&entity.Transaction{})
}

func (repo transactionRepositoryGorm) Save(transaction entity.Transaction) (entity.Transaction, error) {
	err := repo.DB.Create(&transaction).Error
	return transaction, err
}
