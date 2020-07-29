package models

import "github.com/jinzhu/gorm"

type Class struct {
	gorm.Model
	ClassName string `json:"class_name"`
}

func (u Class) TableName() string {
	return "classes"
}

func NewClassModel() *Class {
	return &Class{}
}
func (c *Class) CreateClass(class *Class) (*Class, error) {
	err := DB.Create(class).Error

	return class, err
}

func (c *Class) GetAllClasses() ([]Class, error) {
	var classes []Class

	err := DB.Find(&classes).Error

	return classes, err
}
