package models

import "github.com/jinzhu/gorm"

type UserModel struct {
	gorm.Model
	Username string `json:"user_name"`
	Age      int    `json:"age"`
	ClassID  uint   `json:"class_id"`
	Class    Class  `gorm:"ForeignKey:ClassID" json:"class"`
}

func (u UserModel) TableName() string {
	return "users"
}

func (c *UserModel) CreateUser(user User) (User, error) {
	err := DB.Create(&user).Error

	return user, err
}

func (c *UserModel) GetAllUsers() ([]UserModel, error) {
	var users []UserModel

	err := DB.Preload("Classes").Find(&users).Error

	return users, err
}
