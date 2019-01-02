package app

import (
	"Structure/src/system/router"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type Server struct {
	port string
	DB   *gorm.DB
}

func NewServer() Server {
	return Server{}
}

// init all vals
func (s *Server) Init(port string, db *gorm.DB) {
	log.Println("Initializing server...")
	s.port = ":" + port
	s.DB = db
}

// start the server
func (s *Server) Start() {
	log.Println("Starting server on port" + s.port)

	// initialize routes
	r := mux.NewRouter().StrictSlash(true)
	r.Use(AuthMiddleware)
	router.Routes(r, s.DB)

	http.ListenAndServe(s.port, r)
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Println("You are in : ", r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
