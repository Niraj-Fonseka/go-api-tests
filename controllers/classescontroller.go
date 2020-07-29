package controllers

import (
	"github.com/gin-gonic/gin"
)

type ClassesController struct {
}

func NewClassesController() *ClassesController {
	return &ClassesController{}
}

func (l *ClassesController) GetClasses(c *gin.Context) {

}

func (l *ClassesController) CreateClass(c *gin.Context) {

}
