package models

type ClassModelInterface interface {
	CreateClass(class *Class) (*Class, error)
	GetAllClasses() ([]Class, error)
}
