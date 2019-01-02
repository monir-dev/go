package router

import (
	jwt "Structure/src/system/middleware"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func Routes(r *mux.Router, db *gorm.DB) {
	// r.HandleFunc("/", middleware.AuthRequired(indexGetHandler)).Methods("GET")
	WebRoutes(r, db)

	fs := http.FileServer(http.Dir("./static/"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
}

func IsAuthenticated(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		bearer := r.Header.Get("Authorization")
		if bearer != "" {
			token := strings.Replace(bearer, "Bearer ", "", 1)
			response, err := jwt.PurseToken(token)

			if err != nil {
				w.Write([]byte("err"))
			} else {
				w.Write([]byte(response))
				// endpoint(w, r)
			}

		} else {
			fmt.Fprintf(w, "Not Authorized")
		}
	})
}
