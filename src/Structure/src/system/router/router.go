package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Routes(r *mux.Router) {
	// r.HandleFunc("/", middleware.AuthRequired(indexGetHandler)).Methods("GET")
	WebRoutes(r)

	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
}
