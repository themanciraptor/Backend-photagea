package accountapi

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	accountservice "github.com/themanciraptor/Backend-photagea/internal/account/service"
)

// Interface for the account API
type Interface interface {
	SignIn(http.ResponseWriter, *http.Request)
	Register(http.ResponseWriter, *http.Request)
	Refresh(http.ResponseWriter, *http.Request)
}

// AccountAPI is the API for account related requests
type AccountAPI struct {
	accountService accountservice.Interface
}

type accountRequest struct {
	Email    string `json:"Email"`
	Password string `json:"Password"`
}

// Initialize a new account API
func Initialize(a accountservice.Interface) Interface {
	return &AccountAPI{accountService: a}
}

// SignIn signs in the account and returns a JWT for their use
func (a *AccountAPI) SignIn(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	c := accountRequest{}

	err := d.Decode(&c)
	if err != nil {
		log.Printf("Unable to read request body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	j, err := a.accountService.SignIn(r.Context(), c.Email, c.Password)
	if err != nil {
		log.Printf("Unable to sign in request body: %s", err)
		w.WriteHeader(http.StatusUnauthorized)
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	io.WriteString(w, j)
}

// Register registers a new account
func (a *AccountAPI) Register(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	c := accountRequest{}

	err := d.Decode(&c)
	if err != nil {
		log.Printf("Unable to read request body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = a.accountService.Create(r.Context(), c.Email, c.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusCreated)
}

// Refresh provides a fresh token
func (a *AccountAPI) Refresh(w http.ResponseWriter, r *http.Request) {
	j, err := a.accountService.RefreshToken(r)
	if err != nil {
		log.Printf("Authentication failure: %s", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	io.WriteString(w, j)
}
