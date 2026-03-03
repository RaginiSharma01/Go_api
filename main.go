package main

import (
	"api/config"
	"api/db"
	"api/handler"
	"api/repository"
	"api/router"
	"api/services"
	"log"
	"net/http"
)

func main() {

	dbconnection := config.LoadConfig()

	database, err := db.ConnectDb(dbconnection)
	if err != nil {
		log.Fatal(err)
	}

	userRepo := repository.NewUser(database.Pool)
	userService := services.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	router.SetupRoutes(userHandler)

	http.ListenAndServe(":8080", nil)
}
