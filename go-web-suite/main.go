package main

import (
	"fmt"
	"go-web-suite/config"
	"go-web-suite/models"

	"go-web-suite/routes"
)

func main() {

	config.ConnectDatabases()
	config.DB.AutoMigrate(&models.User{})

	router := routes.SetupRoutes()

	port := "8080"
	fmt.Println("Server running on port", port)
	router.Run(":" + port)
}
