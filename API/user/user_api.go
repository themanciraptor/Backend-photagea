package userapi

import (
	"encoding/json"
	"net/http"

	"github.com/themancirapter/Backend-photagea/Internal/user/userservice"
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
	w.Write(json.Marshal(userservice.Get(12)))
}

// Create a user
func (u *UserAPI) Create(w http.ResponseWriter, r *http.Request) {

}

// Update a user
func (u *UserAPI) Update(w http.ResponseWriter, r *http.Request) {

}
