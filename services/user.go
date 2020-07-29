package services

import (
	"go-api-tests/models"
)


type UserService struct{
	UserModel *models.User
}


func NewUserService() *UserService{
	return &UserService{}
}


func (u *UserService) CreateUser(user *models.User) (*models.User , error){
	return u.UserModel.CreateUser(user)
}

func (u *UserService) GetAllUsers() ([]models.User , error) {
	return u.UserModel.GetAllUsers()
}