package app

import (
	"Structure/src/system/router"
	"log"
	"net/http"

	"github.com/go-xorm/xorm"
	"github.com/gorilla/mux"
)

type Server struct {
	port string
	Db   *xorm.Engine
}

func NewServer() Server {
	return Server{}
}

// init all vals
func (s *Server) Init(port string, db *xorm.Engine) {
	log.Println("Initializing server...")
	s.port = ":" + port
	s.Db = db
}

// start the server
func (s *Server) Start() {
	log.Println("Starting server on port" + s.port)

	// initialize routes
	r := mux.NewRouter().StrictSlash(true)
	r.Use(AuthMiddleware)
	router.Routes(r)

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
