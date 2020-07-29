package controllers

import (
	"fmt"
	"go-api-tests/models"
	"go-api-tests/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService *services.UserService
}

func NewUserController() *UserController {
	return &UserController{}
}

func (u *UserController) CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(500, gin.H{
			"error": fmt.Errorf("Unable to bind the json"),
		})
		return
	}

	_, err := u.UserService.CreateUser(&user)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, nil)
}

func (u *UserController) GetUsers(c *gin.Context) {
	users, err := u.UserService.GetAllUsers()

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"users": users,
	})
}
