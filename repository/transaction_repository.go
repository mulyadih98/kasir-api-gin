package repository

import "kasir-api-gin/domains/entity"

type TransactionRepository interface {
	Migrate() error
	Save(entity.Transaction) (entity.Transaction, error)
	GetById(id string) (entity.Transaction, error)
}
