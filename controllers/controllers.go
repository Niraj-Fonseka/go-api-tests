package controllers

import "go-api-tests/services"

type Controllers struct {
	HealthController  *HealthController
	UserController    *UserController
	ClassesController *ClassesController
}

func InitController(svcs *services.Services) *Controllers {
	var controllers Controllers

	healthContoller := NewHealthController()

	controllers.HealthController = healthContoller

	return &controllers

}
