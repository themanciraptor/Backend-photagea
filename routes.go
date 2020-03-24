package main

import (
	"net/http"

	userapi "github.com/themanciraptor/Backend-photagea/API/user"
)

// RegisterRoutes registers each handler with their respective path
func RegisterRoutes(user userapi.Interface) {
	// routes contains the central register of all routes
	routes := map[string]func(http.ResponseWriter, *http.Request){
		"/user/get":    user.Get,
		"/user/create": user.Create,
		"/user/update": user.Update,
	}

	for path, route := range routes {
		http.HandleFunc(path, route)
	}
}
