package models

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//DBInit initializes the database
func DBInit() (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	dbinfo := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		"postgres",
		"secret",
		"0.0.0.0",
		"5432",
		"app",
	)
	for i := 0; i < 10; i++ {
		db, err = gorm.Open("postgres", dbinfo) // gorm checks Ping on Open
		if err == nil {
			break
		}
		log.Printf("Error while connecting to DB : %s", err.Error())
		log.Println("Trying to connect ... waiting 5 seconds")
		time.Sleep(5 * time.Second)
	}

	if err != nil {
		return db, err
	}

	//create tables at start
	err = db.AutoMigrate(&Class{}, &User{}).Error

	if err != nil {
		log.Fatal(err)
	}

	//add relationship
	err = db.Model(&User{}).AddForeignKey("id", "classes(id)", "CASCADE", "CASCADE").Error

	if err != nil {
		log.Fatal(err)
	}

	return db, err
}
