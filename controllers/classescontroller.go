package controllers

import (
	"fmt"
	"go-api-tests/models"
	"go-api-tests/services"

	"github.com/gin-gonic/gin"
)

type ClassesController struct {
	ClassServiceInterface services.ClassServiceInterface
}

func NewClassesController() *ClassesController {
	return &ClassesController{}
}

func (l *ClassesController) GetClasses(c *gin.Context) {
	classes, err := l.ClassServiceInterface.GetAllClasses()

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"classes": classes,
	})
}

func (l *ClassesController) CreateClass(c *gin.Context) {
	var class models.Class
	if err := c.ShouldBind(&class); err != nil {
		c.JSON(500, gin.H{
			"error": fmt.Errorf("Unable to bind the json"),
		})
		return
	}

	_, err := l.ClassServiceInterface.CreateClass(&class)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, nil)
}
