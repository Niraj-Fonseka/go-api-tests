package services

import (
	"go-api-tests/models"

	"github.com/jinzhu/gorm"
)

type Services struct {
	UserService  *UserService
	ClassService *ClassService
}

func InitServices(db *gorm.DB) *Services {
	userModel := models.NewUserModel(db)
	classModel := models.NewClassModel(db)
	userService := NewUserService()
	userService.UserModelInterface = userModel

	classService := NewClassService()
	classService.ClassModelInterface = classModel

	return &Services{
		UserService:  userService,
		ClassService: classService,
	}
}
