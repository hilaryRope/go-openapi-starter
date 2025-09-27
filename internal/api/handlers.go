package api

import (
	"encoding/json"
	"go-openapi-starter/internal/models"
	"net/http"
	"strconv"
)

type Handlers struct {
	users  []models.User
	nextID int
}

func NewHandlers() *Handlers {
	return &Handlers{
		users: []models.User{
			{ID: "1", Name: "John"},
			{ID: "2", Name: "Claire"},
		},
		nextID: 3,
	}
}

func (h *Handlers) GetHealth(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{"status": "ok"})
}

func (h *Handlers) GetUsers(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(h.users)
}

func (h *Handlers) PostUsers(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Assign a new ID
	user.ID = strconv.Itoa(h.nextID)
	h.nextID++
	h.users = append(h.users, user)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(user)
}

func (h *Handlers) GetUsersId(w http.ResponseWriter, r *http.Request, id string) {
	for _, user := range h.users {
		if user.ID == id {
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(user)
			return
		}
	}
	http.NotFound(w, r)
}

func (h *Handlers) PutUsersId(w http.ResponseWriter, r *http.Request, id string) {
	var updated models.User
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	for i, user := range h.users {
		if user.ID == id {
			updated.ID = id // Ensure ID remains the same
			h.users[i] = updated
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(updated)
			return
		}
	}
	http.NotFound(w, r)
}

func (h *Handlers) DeleteUsersId(w http.ResponseWriter, r *http.Request, id string) {
	for i, user := range h.users {
		if user.ID == id {
			h.users = append(h.users[:i], h.users[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.NotFound(w, r)
}

func (h *Handlers) PatchUsersId(w http.ResponseWriter, r *http.Request, id string) {
	var userIndex = -1
	var existingUser models.User

	for i, u := range h.users {
		if u.ID == id {
			userIndex = i
			existingUser = u
			break
		}
	}

	if userIndex == -1 {
		http.NotFound(w, r)
		return
	}

	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if name, ok := updates["name"].(string); ok && name != "" {
		existingUser.Name = name
	}

	h.users[userIndex] = existingUser
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(existingUser)
}
