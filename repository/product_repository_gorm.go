package repository

import (
	"kasir-api-gin/domains/entity"

	"gorm.io/gorm"
)

type productRepositoryGorm struct {
	DB *gorm.DB
}

func NewProductRepositoryGorm(db *gorm.DB) ProductRepository {
	return productRepositoryGorm{
		DB: db,
	}
}

func (repo productRepositoryGorm) Migrate() error {
	return repo.DB.AutoMigrate(&entity.Product{})
}

func (repo productRepositoryGorm) Save(product entity.Product) (uint, error) {
	err := repo.DB.Create(&product).Error
	return product.ID, err
}

func (repo productRepositoryGorm) GetAll() (products []entity.Product, err error) {
	err = repo.DB.Find(&products).Error
	return products, err
}

func (repo productRepositoryGorm) GetById(id uint) (product entity.Product, err error) {
	err = repo.DB.First(&product, id).Error
	return product, err
}
