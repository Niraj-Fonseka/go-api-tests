package services

import "go-api-tests/models"

type ClassServiceInterface interface {
	CreateClass(class *models.Class) (*models.Class, error)
	GetAllClasses() ([]models.Class, error)
}
