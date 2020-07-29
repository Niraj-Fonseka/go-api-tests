package controllers

import "go-api-tests/services"

type Controllers struct {
	HealthController  *HealthController
	UserController    *UserController
	ClassesController *ClassesController
}

func InitController(svcs *services.Services) *Controllers {
	var controllers Controllers

	userController := NewUserController()
	userController.UserService = svcs.UserService

	classController := NewClassesController()
	classController.ClassService = svcs.ClassService

	healthContoller := NewHealthController()

	controllers.HealthController = healthContoller
	controllers.ClassesController = classController
	controllers.UserController = userController
	return &controllers

}
