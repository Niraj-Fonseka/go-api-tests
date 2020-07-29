package services

import "go-api-tests/models"

type ClassService struct {
	ClassModel *models.Class
}

func NewClassService() *ClassService {
	return &ClassService{}
}

func (c *ClassService) CreateClass(class *models.Class) (*models.Class, error) {
	return c.ClassModel.CreateClass(class)
}

func (c *ClassService) GetAllClasses() ([]models.Class, error) {
	return c.ClassModel.GetAllClasses()
}
