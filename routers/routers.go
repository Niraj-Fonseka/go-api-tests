package routers

import (
	"go-api-tests/controllers"
	"log"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

var (
	port = ":8080"
)

func InitRouters(controllers *controllers.Controllers) {
	router := gin.Default()

	router.Use(cors.Default())

	v1 := router.Group("v1")

	v1.GET("/health", controllers.HealthController.HealthCheck)
	v1.GET("/users", controllers.UserController.GetUsers)
	v1.POST("/users", controllers.UserController.CreateUser)
	v1.GET("/classes", controllers.ClassesController.GetClasses)
	v1.POST("/classes", controllers.ClassesController.CreateClass)

	log.Println("Server running on port : ", port)
	log.Fatal(router.Run(port))
}
