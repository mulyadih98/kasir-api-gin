package repository

import "kasir-api-gin/domains/entity"

type ProductRepository interface {
	Save(entity.Product) (uint, error)
}
