package imageapi

import (
	"encoding/json"
	"log"
	"net/http"

	accountservice "github.com/themanciraptor/Backend-photagea/internal/account/service"
	image "github.com/themanciraptor/Backend-photagea/internal/image/model"
	imageservice "github.com/themanciraptor/Backend-photagea/internal/image/service"
)

// Interface for the image API
type Interface interface {
	Create(http.ResponseWriter, *http.Request)
	List(http.ResponseWriter, *http.Request)
}

//ImageAPI is the API for image related requests
type ImageAPI struct {
	imageService   imageservice.Interface
	accountService accountservice.Interface
}

// imageDataContainer is a temporary container to make json deserialization easier
type imageDataContainer struct {
	URL string `json:"URL"`
}

// Initialize a new instance of the image API
func Initialize(i imageservice.Interface, a accountservice.Interface) Interface {
	return &ImageAPI{imageService: i, accountService: a}
}

// Create a image
func (i *ImageAPI) Create(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	c := imageDataContainer{}

	accountID, err := i.accountService.Verify(r)
	if err != nil {
		log.Printf("Authentication failure: %s", err)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	err = d.Decode(&c)
	if err != nil {
		log.Printf("Unable to read request body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = i.imageService.Create(r.Context(), c.URL, accountID)
	if err != nil {
		log.Printf("Unable to create image for account %d: %s", accountID, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

type listWindow struct {
	Cursor int64 `json:"cursor"`
	Limit  int   `json:"limit"`
}

type listImagesResponse struct {
	listWindow
	Images []*image.Model `json:"images"`
}

// List images attached to account
func (i *ImageAPI) List(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	d.DisallowUnknownFields()

	e := json.NewEncoder(w)

	c := listWindow{}

	accountID, err := i.accountService.Verify(r)
	if err != nil {
		log.Printf("Authentication failed: %s", err)
		w.WriteHeader(http.StatusUnauthorized)
	}

	err = d.Decode(&c)
	if err != nil {
		log.Printf("Unable to read request body: %s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	images, cursor, err := i.imageService.List(r.Context(), accountID, c.Limit, c.Cursor)
	if err != nil {
		log.Printf("Unable to list images for account %d: %s", accountID, err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	e.Encode(&listImagesResponse{
		listWindow: listWindow{Cursor: cursor, Limit: c.Limit},
		Images:     images,
	})
}
