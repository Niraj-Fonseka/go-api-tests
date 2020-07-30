package services

import (
	"go-api-tests/models"
	"testing"

	"gotest.tools/assert"
)

func TestGetAllClassesSuccess(t *testing.T) {

	mockModel := new(models.MockClass)
	classService := NewClassService()
	classService.ClassModelInterface = mockModel

	mockModel.On("GetAllClasses").Return([]models.Class{}, nil)
	_, err := classService.GetAllClasses()

	assert.Equal(t, nil, err, "No error")
	mockModel.AssertExpectations(t)

}

func TestGetAllClassFail(t *testing.T) {
	mockModel := new(models.MockClass)
	classService := NewClassService()
	classService.ClassModelInterface = mockModel

	mockModel.On("GetAllClasses").Return([]models.Class{}, nil)
	_, err := classService.GetAllClasses()

	assert.Equal(t, nil, err, "No error")
	mockModel.AssertExpectations(t)
}
