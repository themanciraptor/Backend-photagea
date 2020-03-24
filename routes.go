package main

import (
	"net/http"

	"Github.com/themancirapter/Backend-photagea/API/userapi"
)

// RegisterRoutes registers each handler with their respective path
func RegisterRoutes(user userapi.Interface) {
	// routes contains the central register of all routes
	routes := map[string]func(http.ResponseWriter, *http.Request){
		"/user/get": user.Get,
	}

	for path, route := range routes {
		http.HandleFunc(path, route)
	}
}
