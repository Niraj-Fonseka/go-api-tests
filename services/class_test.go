package services

import (
	"errors"
	"go-api-tests/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAllClassesSuccess(t *testing.T) {

	mockModel := new(models.MockClass)
	classService := NewClassService()
	classService.ClassModelInterface = mockModel

	mockModel.On("GetAllClasses").Return([]models.Class{}, nil)
	_, err := classService.GetAllClasses()

	assert.Equal(t, nil, err, "There was no error")
	mockModel.AssertExpectations(t)

}

func TestGetAllClassFail(t *testing.T) {
	mockModel := new(models.MockClass)
	classService := NewClassService()
	classService.ClassModelInterface = mockModel
	modelErr := errors.New("Unable to get all classes")
	var classes []models.Class
	mockModel.On("GetAllClasses").Return(classes, modelErr)
	resultClasses, err := classService.GetAllClasses()

	assert.Equal(t, err, err, "Error should be not nil when there's an error in models")
	assert.Equal(t, len(resultClasses), len(classes), "Should Be Equal")
	mockModel.AssertExpectations(t)
}

func TestCreateClass(t *testing.T) {
	mockModel := new(models.MockClass)
	classService := NewClassService()
	classService.ClassModelInterface = mockModel

	class := &models.Class{ClassName: "class_one"}

	var err error
	mockModel.On("CreateClass", class).Return(class, err)
	t.Log(class)

	createdClass, createdErr := classService.CreateClass(class)

	t.Log(createdClass)
	assert.Equal(t, err, createdErr, "Error should be not nil when there's an error in models")

	assert.Equal(t, class, createdClass, "Should Be Equal")
}
