package app

import (
	"log"
	"net/http"
)

type Server struct {
	port string
}

func NewServer() Server {
	return Server{}
}

// init all vals
func (s *Server) Init(port string) {
	log.Println("Initializing server...")
	s.port = ":" + port
}

// start the server
func (s *Server) Start() {
	log.Println("Starting server on port" + s.port)
	http.ListenAndServe(s.port, nil)
}
