package service

import (
	"kasir-api-gin/domains/entity"
	"kasir-api-gin/repository"
	"log"
)

type productService struct {
	productRepository repository.ProductRepository
}

type ProductService interface {
	Save(entity.Product) (uint, error)
	GetAll() ([]entity.Product, error)
	GetById(uint) (entity.Product, error)
}

func NewProductService(productRepo repository.ProductRepository) ProductService {
	if err := productRepo.Migrate(); err != nil {
		log.Panic(err.Error())
	}
	return productService{
		productRepository: productRepo,
	}
}

func (repo productService) Save(product entity.Product) (uint, error) {
	newProduct, err := repo.productRepository.Save(product)
	return newProduct, err
}

func (repo productService) GetAll() ([]entity.Product, error) {
	return repo.productRepository.GetAll()
}

func (repo productService) GetById(id uint) (entity.Product, error) {
	return repo.productRepository.GetById(id)
}

func (repo productService) Edir(id string, product entity.Product) (entity.Product, error) {
	return repo.productRepository.Edit(id, product)
}
