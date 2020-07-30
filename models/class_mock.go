package models

import "github.com/stretchr/testify/mock"

type MockClass struct {
	mock.Mock
}

func (c *MockClass) GetAllClasses() ([]Class, error) {
	args := c.Called()
	return args.Get(0).([]Class), args.Error(1)
}

func (c *MockClass) CreateClass(class *Class) (*Class, error) {
	args := c.Called(class)
	return args.Get(0).(*Class), args.Error(1)
}
