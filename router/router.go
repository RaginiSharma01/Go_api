package router

import (
	"api/handler"
	"net/http"
)

func SetupRoutes(userHandler *handler.UserHandler) {

	http.HandleFunc("/api/v1/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			userHandler.CreateUser(w, r)
		case http.MethodGet:
			userHandler.GetUsers(w, r)
		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/api/v1/users/", func(w http.ResponseWriter, r *http.Request) {

		switch r.Method {
		case http.MethodGet:
			userHandler.GetUser(w, r)

		case http.MethodDelete:
			userHandler.DeleteUser(w, r)

		case http.MethodPut:
			userHandler.UpdateUser(w, r)

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}

	})
}
