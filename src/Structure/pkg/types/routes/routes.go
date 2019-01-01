package routes

import (
	"net/http"
)

type Routes []Route

type Route struct {
	Name    string
	Method  string
	Pattern string
	// HandlerFunc http.HandleFunc
}

type SubRoutePackage struct {
	Routes     Routes
	Middleware func(next http.Handler) http.Handler
}
