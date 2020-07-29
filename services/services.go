package services

import "go-api-tests/models"

type Services struct {
	UserService  *UserService
	ClassService *ClassService
}

func InitServices() *Services {
	userModel := models.NewUserModel()
	classModel := models.NewClassModel()
	userService := NewUserService()
	userService.UserModel = userModel

	classService := NewClassService()
	classService.ClassModel = classModel

	return &Services{
		UserService:  userService,
		ClassService: classService,
	}
}
