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
func (repo transactionRepositoryGorm) GetById(id string) (transaction entity.Transaction, err error) {
	err = repo.DB.Where("id = ?", id).Preload("TransactionDetail.Product").Preload("TransactionDetail").Find(&transaction).Error
	return transaction, err
}
