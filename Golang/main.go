package main

import (
	"Jur/config"
	"Jur/controller"
	"Jur/entities"

	"Jur/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// Main Golang
	rute := gin.Default()

	database, err := config.GalaSetup()
	if err != nil {
		log.Fatalf("Fail : %v", err)
		return
	}

	errto := database.AutoMigrate(&entities.SampleTabler{})

	if errto != nil {
		log.Fatalf("Migration Fail : %v", errto)
	}

	rute.GET("/", controller.MainHallo)

	router.SampleRouter(rute)

	if err := rute.Run(); err != nil {
		log.Fatalf("Fail to Start Server : %v", err)
	}
}
