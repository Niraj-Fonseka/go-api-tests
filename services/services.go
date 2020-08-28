package services

import (
	"go-api-tests/models"

	"github.com/jinzhu/gorm"
)

type Services struct {
	UserService  UserServiceInterface
	ClassService ClassServiceInterface
}

func InitServices(db *gorm.DB) *Services {
	userModel := models.NewUserModel(db)
	classModel := models.NewClassModel(db)
	userService := NewUserService()
	userService.UserModel = userModel

	classService := NewClassService()
	classService.ClassModel = classModel

	return &Services{
		UserService:  userService,
		ClassService: classService,
	}
}
