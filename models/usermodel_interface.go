package models

type UserModelInterface interface {
	CreateUser(user *User) (*User, error)
	GetAllUsers() ([]User, error)
}
