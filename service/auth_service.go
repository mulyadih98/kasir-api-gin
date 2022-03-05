package service

import (
	"errors"
	"kasir-api-gin/domains/entity"
	"kasir-api-gin/helper"
	"kasir-api-gin/repository"
	"log"
	"strconv"
)

type authService struct {
	authRepository repository.AuthRepository
	userRepository repository.UserRepository
	passwordHash   helper.PasswordHash
	tokenize       helper.TokenJWT
}

type AuthService interface {
	Register(entity.User) (uint, error)
	Login(entity.LoginInput) (entity.Auth, error)
	Refresh(string) (string, error)
	Logout(string) error
}

func NewAuthService(userRepo repository.UserRepository,
	hash helper.PasswordHash,
	token helper.TokenJWT,
	auth repository.AuthRepository,
) AuthService {
	if err := userRepo.Migrate(); err != nil {
		log.Panic(err.Error())
	}
	if err := auth.Migrate(); err != nil {
		log.Panic(err.Error())
	}
	return authService{
		userRepository: userRepo,
		passwordHash:   hash,
		tokenize:       token,
		authRepository: auth,
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

func (service authService) Login(input entity.LoginInput) (token entity.Auth, err error) {
	userLogin, err := service.userRepository.GetByEmail(input.Email)
	if err != nil {
		return token, errors.New("email tidak terdaftar")
	}

	if err := service.passwordHash.Compare(userLogin.Password, input.Password); err != nil {
		return token, errors.New("password salah")
	}

	var tokenString string
	tokenString, _ = service.tokenize.Generate(userLogin.ID)
	refresToken, _ := service.tokenize.GenerateRefresh(userLogin.ID)
	service.authRepository.Save(entity.AuthToken{RefreshToken: refresToken, UserID: userLogin.ID})
	token = entity.Auth{
		Token:        tokenString,
		RefreshToken: refresToken,
	}
	return token, err
}

func (service authService) Refresh(refreshToken string) (string, error) {
	id, err := service.authRepository.GetByToken(refreshToken)
	if err != nil {
		return "", err
	}
	userId, _ := strconv.ParseUint(id, 10, 32)
	return service.tokenize.Generate(uint(userId))
}

func (service authService) Logout(user_id string) error {
	err := service.authRepository.Delete(user_id)
	return err
}
