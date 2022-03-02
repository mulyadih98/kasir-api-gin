package repository

import (
	"kasir-api-gin/domains/entity"

	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return userRepository{
		DB: db,
	}
}

func (repo userRepository) Migrate() error {
	return repo.DB.AutoMigrate(&entity.User{})
}

func (repo userRepository) Save(newUser entity.User) (uint, error) {
	err := repo.DB.Create(&newUser).Error
	return newUser.ID, err
}

func (repo userRepository) Edit(id uint, newUser entity.User) (user entity.User, err error) {
	repo.DB.First(&user, id)
	if newUser.Username != "" {
		user.Username = newUser.Username
	}

	if newUser.Email != "" {
		user.Email = newUser.Email
	}

	if newUser.Password != "" {
		user.Password = newUser.Password
	}

	repo.DB.Save(&user)
	return user, err
}

func (repo userRepository) GetById(id uint) (user entity.User, err error) {
	err = repo.DB.First(&user, id).Error
	return user, err
}

func (repo userRepository) GetByEmail(email string) (user entity.User, err error) {
	err = repo.DB.Where("email = ? ", email).First(&user).Error
	return user, err
}

func (repo userRepository) GetByUsername(username string) (user entity.User, err error) {
	err = repo.DB.Where("username = ? ", username).First(&user).Error
	return user, err
}

func (repo userRepository) Delete(id uint) error {
	var user entity.User
	err := repo.DB.Delete(&user, id).Error
	return err
}
