package main

import (
	"net/http"

	accountapi "github.com/themanciraptor/Backend-photagea/API/account"
	imageapi "github.com/themanciraptor/Backend-photagea/API/image"
	userapi "github.com/themanciraptor/Backend-photagea/API/user"
)

// RegisterRoutes registers each handler with their respective path
func RegisterRoutes(user userapi.Interface, accounts accountapi.Interface, images imageapi.Interface) {
	// routes contains the central register of all routes
	routes := map[string]func(http.ResponseWriter, *http.Request){
		"/user/get":         user.Get,
		"/user/create":      user.Create,
		"/user/update":      user.Update,
		"/account/signin":   accounts.SignIn,
		"/account/register": accounts.Register,
		"/images/create":    images.Create,
		"/images/list":      images.List,
	}

	for path, route := range routes {
		http.HandleFunc(path, route)
	}
}
