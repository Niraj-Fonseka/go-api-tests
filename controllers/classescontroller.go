package controllers

import (
	"fmt"
	"go-api-tests/models"
	"go-api-tests/services"

	"github.com/gin-gonic/gin"
)

type ClassesController struct {
	ClassService *services.ClassService
}

func NewClassesController() *ClassesController {
	return &ClassesController{}
}

func (l *ClassesController) GetClasses(c *gin.Context) {
	classes, err := l.ClassService.GetAllClasses()

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

	_, err := l.ClassService.CreateClass(&class)

	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, nil)
}
