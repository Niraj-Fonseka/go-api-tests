package services

import (
	"go-api-tests/models"
)

type UserService struct {
	UserModelInterface models.UserModelInterface
}

func NewUserService() *UserService {
	return &UserService{}
}

func (u *UserService) CreateUser(user *models.User) (*models.User, error) {
	return u.UserModelInterface.CreateUser(user)
}

func (u *UserService) GetAllUsers() ([]models.User, error) {
	return u.UserModelInterface.GetAllUsers()
}
