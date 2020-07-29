package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `json:"user_name"`
	Age      int    `json:"age"`
	ClassID  uint   `json:"class_id"`
	Class    Class  `gorm:"ForeignKey:ClassID" json:"class"`
}

func (u User) TableName() string {
	return "users"
}

func NewUserModel() *User {
	return &User{}
}

func (c *User) CreateUser(user *User) (*User, error) {
	err := DB.Create(user).Error

	return user, err
}

func (c *User) GetAllUsers() ([]User, error) {
	var users []User

	err := DB.Preload("Class").Find(&users).Error

	return users, err
}
