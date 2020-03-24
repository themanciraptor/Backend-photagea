package userapi

import (
	"encoding/json"
	"log"
	"net/http"

	userservice "github.com/themanciraptor/Backend-photagea/internal/user/service"
)

// Interface for the user API
type Interface interface {
	Get(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
}

// UserAPI is the API for user related requests
type UserAPI struct {
	userService userservice.Interface
}

// Initialize a new instance of the user API
func Initialize(u userservice.Interface) Interface {
	return &UserAPI{userService: u}
}

// Get a user's details
func (u *UserAPI) Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user, err := u.userService.Get(r.Context(), 12)
	if err != nil {
		log.Printf("Unable to find user: %d", 12)
	}

	encodedUser, err := json.Marshal(user)
	if err != nil {
		log.Fatalf("Unable to serialize user: %d", 12)
	}

	w.Write(encodedUser)
}

// Create a user
func (u *UserAPI) Create(w http.ResponseWriter, r *http.Request) {

}

// Update a user
func (u *UserAPI) Update(w http.ResponseWriter, r *http.Request) {

}
