package services

import "go-api-tests/models"

type UserServiceInterface interface {
	CreateUser(user *models.User) (*models.User, error)
	GetAllUsers() ([]models.User, error)
}
