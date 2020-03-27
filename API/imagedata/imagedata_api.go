package imagedataapi

import (
	"encoding/json"
	"io/ioutil"
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
	imageDataService imagedataservice.Interface
	accountService   accountservice.Interface
}

// Initialize a new instance of the user API
func Initialize(i imagedataservice.Interface, a accountservice.Interface) Interface {
	return &imageDataAPI{imageDataService: i, accountService: a}
}

func (i *imageDataAPI) Get(w http.ResponseWriter, r *http.Request) {
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

	id, err := i.imageDataService.Get(r.Context(), accountID, int64(idid))
	if err != nil {
		log.Println("Failed to retrieve image with id: ", idid)
		log.Println("Error: ", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", id.MimeType)
	w.Write(id.ImageData)
}

type uploadResponse struct {
	URL string
}

func (i *imageDataAPI) Upload(w http.ResponseWriter, r *http.Request) {
	e := json.NewEncoder(w)

	accountID, err := i.accountService.Verify(r)
	if err != nil {
		log.Printf("Authentication failed: %s", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	mimetype := r.Header.Get("Content-Type")

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("Unable to read and store image")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	url, err := i.imageDataService.Upload(r.Context(), accountID, mimetype, data)
	if err != nil {
		log.Println("Unable to save image: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	e.Encode(&uploadResponse{URL: url})
}
