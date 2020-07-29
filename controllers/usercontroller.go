package controllers

import "github.com/gin-gonic/gin"

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (u *UserController) CreateUser(c *gin.Context) {

}

func (u *UserController) GetUsers(c *gin.Context) {

}
