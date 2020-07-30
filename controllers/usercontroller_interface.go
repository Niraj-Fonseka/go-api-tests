package controllers 


import(
	"github.com/gin-gonic/gin"
)
type UserControllerInterface interface {
	CreateUser(c *gin.Context) 
	GetUsers(c *gin.Context)
}