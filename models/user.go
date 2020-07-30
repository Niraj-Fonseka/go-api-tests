package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `json:"user_name"`
	Age      int    `json:"age"`
	ClassID  uint   `json:"class_id"`
	Class    Class  `gorm:"ForeignKey:ClassID" json:"class"`
}

type UserModel struct {
	db *gorm.DB
}

func (u User) TableName() string {
	return "users"
}

func NewUserModel(db *gorm.DB) *UserModel {
	return &UserModel{
		db: db,
	}
}

func (c *UserModel) CreateUser(user *User) (*User, error) {
	err := c.db.Create(user).Error

	return user, err
}

func (c *UserModel) GetAllUsers() ([]User, error) {
	var users []User

	err := c.db.Preload("Class").Find(&users).Error

	return users, err
}
