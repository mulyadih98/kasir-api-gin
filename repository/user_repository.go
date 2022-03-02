package repository

import "kasir-api-gin/domains/entity"

type UserRepository interface {
	Save(entity.User) (uint, error)
	Edit(id uint, user entity.User) (entity.User, error)
	GetById(id uint) (entity.User, error)
	GetByEmail(email string) (entity.User, error)
	GetByUsername(email string) (entity.User, error)
	Delete(id uint) error
	Migrate() error
}
