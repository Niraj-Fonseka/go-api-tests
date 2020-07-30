package controllers

import "github.com/gin-gonic/gin"

type ClassControllerInterface interface {
	GetClasses(c *gin.Context)
	CreateClass(c *gin.Context)
}
