package repository

import "kasir-api-gin/domains/entity"

type ProductRepository interface {
	Migrate() error
	Save(entity.Product) (uint, error)
	GetAll() ([]entity.Product, error)
	GetById(id uint) (entity.Product, error)
	Edit(string, entity.Product) (entity.Product, error)
	Delete(id string) error
	ReduceStock(id string, qty int) error
}
