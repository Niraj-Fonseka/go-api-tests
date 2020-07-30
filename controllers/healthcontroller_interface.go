package controllers

import "github.com/gin-gonic/gin"

type HealthControllerInterface interface {
	HealthCheck(c *gin.Context)
}
