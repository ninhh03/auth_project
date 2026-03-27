package services

import (
	"auth_project/internal/models"
	"auth_project/internal/repository"
	"auth_project/pkg/utils"
	"errors"
)

type authService struct {
	ur repository.UserRepository
}

func NewAuthService(ur repository.UserRepository) *authService {
	return &authService{ur: ur}
}

func (as *authService) Register(email string, password string) error {
	existingUser, err := as.ur.GetUserByEmail(email)
	if err != nil {
		return err
	}
	if existingUser != nil {
		return errors.New("email already exists")
	}
	hashPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}
	user := &models.User{
		Email:    email,
		Password: string(hashPassword),
	}
	return as.ur.CreateUser(user)
}

func (as *authService) Login(email string, password string) error {
	existingUser, err := as.ur.GetUserByEmail(email)
	if err != nil {
		return err
	}
	if existingUser == nil {
		return errors.New("user don't exist")
	}

	if err := utils.CheckPassword(password, existingUser.Password); err != nil {
		return errors.New("email or password incorect!")
	}
	return nil
}
