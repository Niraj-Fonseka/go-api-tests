package services

import (
	"go-api-tests/models"

	"github.com/stretchr/testify/mock"
)

type MockClassService struct {
	mock.Mock
}

func (c *MockClassService) GetAllClasses() ([]models.Class, error) {
	args := c.Called()
	return args.Get(0).([]models.Class), args.Error(1)
}

func (c *MockClassService) CreateClass(class *models.Class) (*models.Class, error) {
	args := c.Called(class)
	return args.Get(0).(*models.Class), args.Error(1)
}
