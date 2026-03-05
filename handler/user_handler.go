package handler

import (
	"api/models"
	"api/services"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type UserHandler struct {
	Service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{
		Service: service,
	}
}

func writeJSONError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(map[string]string{
		"error": message,
	})
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		writeJSONError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	defer r.Body.Close()

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid JSON payload")
		return
	}

	// validation
	if user.Name == "" || user.Email == "" {
		writeJSONError(w, http.StatusBadRequest, "name and email are required")
		return
	}

	err = h.Service.CreateUser(user)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "user created successfully",
	})
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		writeJSONError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	users, err := h.Service.GetUsers()
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, "failed to fetch users")
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(users)
}
func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {

	idStr := strings.TrimPrefix(r.URL.Path, "/api/v1/users/")

	if idStr == "" {
		writeJSONError(w, http.StatusBadRequest, "id is required")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid user id")
		return
	}

	user, err := h.Service.GetUserByID(id)
	if err != nil {
		writeJSONError(w, http.StatusNotFound, "user not found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		writeJSONError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/users/")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid user id")
		return
	}

	err = h.Service.DeleteUser(id)
	if err != nil {
		writeJSONError(w, http.StatusNotFound, err.Error())
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		writeJSONError(w, http.StatusMethodNotAllowed, "method not allowed")
		return
	}

	idStr := strings.TrimPrefix(r.URL.Path, "/users/")

	_, err := strconv.Atoi(idStr)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid user id")
		return
	}

	defer r.Body.Close()

	var user models.User

	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid JSON payload")
		return
	}

	if user.Name == "" || user.Email == "" {
		writeJSONError(w, http.StatusBadRequest, "name and email are required")
		return
	}

	// err = h.Service.UpdateUser(id, user)
	// if err != nil {
	// 	writeJSONError(w, http.StatusInternalServerError, err.Error())
	// 	return
	// }

	// w.Header().Set("Content-Type", "application/json")

	// json.NewEncoder(w).Encode(map[string]string{
	// 	"message": "user updated successfully",
	// })
}
