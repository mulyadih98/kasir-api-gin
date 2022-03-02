package service

import (
	"errors"
	"kasir-api-gin/domains/entity"
	"kasir-api-gin/helper"
	"kasir-api-gin/repository"
)

type authService struct {
	userRepository repository.UserRepository
	passwordHash   helper.PasswordHash
	tokenize       helper.TokenJWT
}

type AuthService interface {
	Register(entity.User) (uint, error)
	Login(entity.LoginInput) (string, error)
}

func NewAuthService(userRepo repository.UserRepository, hash helper.PasswordHash, token helper.TokenJWT) AuthService {
	return authService{
		userRepository: userRepo,
		passwordHash:   hash,
		tokenize:       token,
	}
}

func (service authService) Register(user entity.User) (uint, error) {
	password, err := service.passwordHash.Hash(user.Password)
	if err != nil {
		return 0, err
	}

	if user, _ := service.userRepository.GetByEmail(user.Email); user.ID != 0 {
		return 0, errors.New("email telah tigunakan")
	}

	if user, _ := service.userRepository.GetByUsername(user.Username); user.ID != 0 {
		return 0, errors.New("username telah tigunakan")
	}

	id, err := service.userRepository.Save(entity.User{
		Username: user.Username,
		Password: password,
		Email:    user.Email,
	})

	return id, err
}

func (service authService) Login(input entity.LoginInput) (token string, err error) {
	userLogin, err := service.userRepository.GetByEmail(input.Email)
	if err != nil {
		return "", errors.New("email tidak terdaftar")
	}

	if err := service.passwordHash.Compare(userLogin.Password, input.Password); err != nil {
		return "", errors.New("password salah")
	}
	token, err = service.tokenize.Generate(userLogin.ID)

	return token, err
}
