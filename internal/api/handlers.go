package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"go-openapi-starter/internal/models"
	"net/http"
)

var users = []models.User{
	{ID: "1", Name: "Alice"},
	{ID: "2", Name: "Bob"},
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(`{"status":"ok"}`))
	if err != nil {
		return
	}
}

func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(users)
	if err != nil {
		return
	}
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var u models.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	users = append(users, u)
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(u)
	if err != nil {
		return
	}
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	for _, u := range users {
		if u.ID == id {
			err := json.NewEncoder(w).Encode(u)
			if err != nil {
				return
			}
			return
		}
	}
	http.NotFound(w, r)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	for i, u := range users {
		if u.ID == id {
			var updated models.User
			if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			users[i] = updated
			err := json.NewEncoder(w).Encode(updated)
			if err != nil {
				return
			}
			return
		}
	}
	http.NotFound(w, r)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	for i, u := range users {
		if u.ID == mux.Vars(r)["id"] {
			users = append(users[:i], users[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.NotFound(w, r)
}
