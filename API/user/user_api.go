package userapi

import (
	"encoding/json"
	"log"
	"net/http"

	accountservice "github.com/themanciraptor/Backend-photagea/internal/account/service"
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
	userService    userservice.Interface
	accountService accountservice.Interface
}

// userDataContainer is a temporary container to make json deserialization easier
type userDataContainer struct {
	Alias     string `json:"Alias"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
}

// Initialize a new instance of the user API
func Initialize(u userservice.Interface, a accountservice.Interface) Interface {
	return &UserAPI{userService: u, accountService: a}
}

// Get a user's details
func (u *UserAPI) Get(w http.ResponseWriter, r *http.Request) {
	e := json.NewEncoder(w)

	accountID, err := u.accountService.Verify(r)
	if err != nil {
		log.Println("Authentication failure: ", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	user, err := u.userService.Get(r.Context(), accountID)
	if err != nil {
		log.Println("Unable to find user for account: ", accountID)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = e.Encode(user)
	if err != nil {
		log.Fatalf("Unable to serialize user: %d", accountID)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

// Create a user
func (u *UserAPI) Create(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	c := userDataContainer{}

	accountID, err := u.accountService.Verify(r)
	if err != nil {
		log.Printf("A failed attemp at signing in was made: %s", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = d.Decode(&c)
	if err != nil {
		log.Printf("Unable to read request body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = u.userService.Create(r.Context(), accountID, c.Alias, c.FirstName, c.LastName)
	if err != nil {
		log.Printf("Unable to create user for account %d: %s", accountID, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Update a user
func (u *UserAPI) Update(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	c := userDataContainer{}

	accountID, err := u.accountService.Verify(r)
	if err != nil {
		log.Println("A failed attemp at signing in was made: ", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = d.Decode(&c)
	if err != nil {
		log.Println("Unable to read request body: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// TODO: this isn't technically always a 500
	err = u.userService.Update(r.Context(), accountID, c.Alias, c.FirstName, c.LastName)
	if err != nil {
		log.Printf("Unable to update user for account %d: %s", accountID, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
