package main

import (
	"go-api-tests/controllers"
	"go-api-tests/models"
	"go-api-tests/routers"
	"go-api-tests/services"
	"log"
)

func main() {
	log.Println("Hello !")

	_, err := models.DBInit()

	if err != nil {
		log.Fatal(err)
	}
	svcs := services.InitServices()

	ctls := controllers.InitController(svcs)

	routers.InitRouters(ctls)
}
