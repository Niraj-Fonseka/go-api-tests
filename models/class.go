package models

import (
	"github.com/jinzhu/gorm"
)

type Class struct {
	gorm.Model
	ClassName string `json:"class_name"`
	ClassID   int
}

type ClassModel struct {
	db *gorm.DB
}

func (u Class) TableName() string {
	return "classes"
}

func NewClassModel(db *gorm.DB) *ClassModel {
	return &ClassModel{
		db: db,
	}
}
func (c *ClassModel) CreateClass(class *Class) (*Class, error) {

	err := c.db.Create(class).Error

	return class, err
}

func (c *ClassModel) GetAllClasses() ([]Class, error) {
	var classes []Class

	err := c.db.Find(&classes).Error

	return classes, err
}
