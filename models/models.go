package models

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB
var err error

//DBInit initializes the database
func DBInit() (*gorm.DB, error) {

	dbinfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		"postgres",
		"secret",
		"0.0.0.0",
		"5432",
		"app",
	)
	for i := 0; i < 10; i++ {
		DB, err = gorm.Open("postgres", dbinfo) // gorm checks Ping on Open
		if err == nil {
			break
		}
		log.Println("Error while connecting to DB : %s", err.Error())
		log.Println("Trying to connect ... waiting 5 seconds")
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		return DB, err
	}

	//create tables at start
	err := DB.AutoMigrate(&Class{}, &User{}).Error

	if err != nil {
		log.Fatal(err)
	}

	//add relationship
	err = DB.Model(&User{}).AddForeignKey("id", "classes(id)", "CASCADE", "CASCADE").Error

	if err != nil {
		log.Fatal(err)
	}

	return DB, err
}
