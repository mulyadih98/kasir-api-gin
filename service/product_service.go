package service

import (
	"kasir-api-gin/domains/entity"
	"kasir-api-gin/repository"
)

type productService struct {
	productRepository repository.ProductRepository
}

type ProductService interface {
	Save(entity.Product) (uint, error)
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
	productRepo.Migrate()
	return productService{
		productRepository: productRepo,
	}
}

func (repo productService) Save(product entity.Product) (uint, error) {
	newProduct, err := repo.productRepository.Save(product)
	return newProduct, err
}
