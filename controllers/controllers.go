package controllers

import "go-api-tests/services"

type Controllers struct {
	HealthController  HealthControllerInterface
	UserController    UserControllerInterface
	ClassesController ClassControllerInterface
}

func InitController(svcs *services.Services) *Controllers {
	var controllers Controllers

	userController := NewUserController()
	userController.UserServiceInterface = svcs.UserService

	classController := NewClassesController()
	classController.ClassServiceInterface = svcs.ClassService

	healthContoller := NewHealthController()

	controllers.HealthController = healthContoller
	controllers.ClassesController = classController
	controllers.UserController = userController
	return &controllers

}
