package repository

import (
	"kasir-api-gin/domains/entity"
	"log"

	"gorm.io/gorm"
)

type authRepositoryGorm struct {
	DB *gorm.DB
}

type AuthRepository interface {
	Migrate() error
	Save(token entity.AuthToken) error
	GetByToken(token string) (string, error)
	Delete(string) error
}

func NewAuthRepositoryGorm(db *gorm.DB) AuthRepository {
	return authRepositoryGorm{
		DB: db,
	}
}

func (repo authRepositoryGorm) Migrate() error {
	return repo.DB.AutoMigrate(&entity.AuthToken{})
}

func (repo authRepositoryGorm) Save(token entity.AuthToken) error {
	return repo.DB.Create(&token).Error
}

func (repo authRepositoryGorm) GetByToken(token string) (string, error) {
	var refreshToken entity.AuthToken
	err := repo.DB.Where("refresh_token = ? ", token).First(&refreshToken).Error
	return refreshToken.RefreshToken, err
}

func (repo authRepositoryGorm) Delete(user_id string) error {
	var logout entity.AuthToken
	log.Println(user_id)
	err := repo.DB.Unscoped().Where("user_id = ?", user_id).Delete(&logout).Error
	return err
}
