package router

import (
	"api/handler"
	"net/http"
)

func SetupRoutes(userHandler *handler.UserHandler) {

	http.HandleFunc("/users", userHandler.CreateUser)
	http.HandleFunc("/users/id", userHandler.GetUsers)
	http.HandleFunc("/user", userHandler.GetUser)
	http.HandleFunc("/users/", userHandler.DeleteUser)

}
