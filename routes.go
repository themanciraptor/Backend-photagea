package main

import (
	"net/http"

	"github.com/gorilla/mux"
	accountapi "github.com/themanciraptor/Backend-photagea/API/account"
	imageapi "github.com/themanciraptor/Backend-photagea/API/image"
	imagedataapi "github.com/themanciraptor/Backend-photagea/API/imagedata"
	userapi "github.com/themanciraptor/Backend-photagea/API/user"
)

// RegisterRoutes registers each handler with their respective path
func RegisterRoutes(r *mux.Router, user userapi.Interface, accounts accountapi.Interface, images imageapi.Interface, imagedata imagedataapi.Interface) {
	// User Endpoints
	r.HandleFunc("/api/user/get", user.Get).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/api/user/create", user.Create).Methods(http.MethodPut)
	r.HandleFunc("/api/user/update", user.Update).Methods(http.MethodPut)

	// Account Endpoints
	r.HandleFunc("/api/account/signin", accounts.SignIn).Methods(http.MethodPost)
	r.HandleFunc("/api/account/register", accounts.Register).Methods(http.MethodPut)
	r.HandleFunc("/api/account/refresh", accounts.Refresh).Methods(http.MethodPost)

	// Image Endpoints
	r.HandleFunc("/api/images/create", images.Create).Methods(http.MethodPut)
	r.HandleFunc("/api/images/list", images.List).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/api/images/upload", imagedata.Upload).Methods(http.MethodPut)
	r.HandleFunc("/api/images/get", imagedata.Get).Methods(http.MethodGet, http.MethodPost)

	r.HandleFunc("/api/TEST", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusGone) })
}
