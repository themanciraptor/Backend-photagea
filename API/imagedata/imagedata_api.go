package imagedataapi

import (
	"log"
	"net/http"
	"strconv"

	accountservice "github.com/themanciraptor/Backend-photagea/internal/account/service"
	imagedataservice "github.com/themanciraptor/Backend-photagea/internal/imagedata/service"
)

// Interface for the user API
type Interface interface {
	Get(http.ResponseWriter, *http.Request)
	Upload(http.ResponseWriter, *http.Request)
}

// imageDataAPI is the API for user related requests
type imageDataAPI struct {
	imagedataService imagedataservice.Interface
	accountService   accountservice.Interface
}

// Initialize a new instance of the user API
func Initialize(i imagedataservice.Interface, a accountservice.Interface) Interface {
	return &imageDataAPI{imagedataService: i, accountService: a}
}

func (i imageDataAPI) Get(w http.ResponseWriter, r *http.Request) {
	accountID, err := i.accountService.Verify(r)
	if err != nil {
		log.Printf("Authentication failed: %s", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	params := r.URL.Query()

	idid, err := strconv.Atoi(params.Get("id"))
	if err != nil {
		log.Println("query param id missing from request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := i.imagedataService.Get(r.Context(), accountID, int64(idid))
	if err != nil {
		log.Println("Failed to retrieve image with id: ", idid)
		log.Println("Error: ", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", id.MimeType)
	w.Write(id.ImageData)
}

func (i imageDataAPI) Upload(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
