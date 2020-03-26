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
	// routes contains the central register of all routes
	r.HandleFunc("/user/get", user.Get).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/user/create", user.Create).Methods(http.MethodPut)
	r.HandleFunc("/user/update", user.Update).Methods(http.MethodPut)
	r.HandleFunc("/account/signin", accounts.SignIn).Methods(http.MethodPost)
	r.HandleFunc("/account/register", accounts.Register).Methods(http.MethodPut)
	r.HandleFunc("/images/create", images.Create).Methods(http.MethodPut)
	r.HandleFunc("/images/list", images.List).Methods(http.MethodGet, http.MethodPost)
	r.HandleFunc("/images/upload", imagedata.Upload).Methods(http.MethodPut)
	r.HandleFunc("/images/get", imagedata.Get).Methods(http.MethodGet, http.MethodPost)
}
