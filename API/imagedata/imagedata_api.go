package imagedataapi

import (
	"net/http"

	accountservice "github.com/themanciraptor/Backend-photagea/internal/account/service"
)

// Interface for the user API
type Interface interface {
	Get(http.ResponseWriter, *http.Request)
	Upload(http.ResponseWriter, *http.Request)
}

// imageDataAPI is the API for user related requests
type imageDataAPI struct {
	// userService    userservice.Interface
	accountService accountservice.Interface
}

// Initialize a new instance of the user API
func Initialize( /*u userservice.Interface, */ a accountservice.Interface) Interface {
	return &imageDataAPI{ /*userService: u, */ accountService: a}
}

func (i imageDataAPI) Get(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func (i imageDataAPI) Upload(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
