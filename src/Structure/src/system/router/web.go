package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

func WebRoutes(r *mux.Router) {
	// r.HandleFunc("/", middleware.AuthRequired(indexGetHandler)).Methods("GET")
	r.HandleFunc("/", HomeHandler)

	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}
